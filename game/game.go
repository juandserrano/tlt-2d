package game

import (
	"embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// var assetsFS embed.FS

type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StatePause
	StateGameOver
	StateWorldEditor
)

type TurnState int

const (
	TurnPlayer TurnState = iota
	TurnResolving
	TurnComputer
)

type Game struct {
	assets       embed.FS
	Config       GameConfig
	State        GameState
	Turn         TurnState
	Round        Round
	camera       rl.Camera3D
	tiles        map[TileType]Tile
	levels       map[int]Level
	currentLevel int
	debugLevel   uint8
	enemyBag     EnemyBag
	playerHand   Hand
	deck         Deck
	// cardsToPlay    []*Card
	discardPile    Deck
	playerCastle   Castle
	shaders        map[ShaderName]*rl.Shader
	plainTileModel rl.Model
	waterTileModel rl.Model
	enemyModels    map[EnemyType]*rl.Model
	// cardModels     map[CardType]*rl.Model
	cardTextures             map[CardType]*rl.Texture2D
	sunLight                 Light
	spotLight                Light
	frameCount               int
	UI                       UI
	CameraShakeIntensity     float32
	enemyMoveIndex           int
	waitingForMoveAnimation  bool
	waitingForSpawnAnimation bool
	turnTransitionTimer      float32
	uiAlpha                  float32
	endingTurn               bool
	sounds                   map[string]rl.Sound
	music                    map[string]rl.Music
	cardAnimations           []*CardAnimation
	cardSlideAnimations      []*CardSlideAnimation
	ParticleManager          *ParticleManager
}

type CardAnimation struct {
	Card          *Card
	StartPosition rl.Vector2
	TargetEnemy   *Enemy
	Progress      float32
	OnFinish      func()
}

type CardSlideAnimation struct {
	Card           *Card
	StartPosition  rl.Vector2
	TargetPosition rl.Vector2
	Progress       float32
}

func Run(embedFS *embed.FS) {
	game := &Game{}
	game.init(embedFS)
	defer rl.CloseAudioDevice()
	defer rl.CloseWindow()
	game.Loop()
}

func (g *Game) init(embedFS *embed.FS) {
	g.assets = *embedFS
	g.CheckAndLoadConfig(true)
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagMsaa4xHint)
	rl.SetTargetFPS(g.Config.Window.TargetFPS)
	rl.InitWindow(g.Config.Window.Width, g.Config.Window.Height, g.Config.GameName)
	rl.InitAudioDevice()
	g.debugLevel = 0
	g.initCamera()
	g.levels = make(map[int]Level)
	g.tiles = make(map[TileType]Tile)
	g.shaders = make(map[ShaderName]*rl.Shader)
	g.enemyModels = make(map[EnemyType]*rl.Model)
	g.cardTextures = make(map[CardType]*rl.Texture2D)
	g.UI.buttons = make(map[string]*Button)
	g.sounds = make(map[string]rl.Sound)
	g.music = make(map[string]rl.Music)
	g.LoadResources()
	g.initShadersAndLights()
	g.Round = g.NewRound()
	g.State = StatePlaying
}

func (g *Game) NextTurn() {
	g.UpdateEnemies()
}
