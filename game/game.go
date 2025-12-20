package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	Config         GameConfig
	State          GameState
	Turn           TurnState
	Round          Round
	camera         rl.Camera3D
	tiles          map[TileType]Tile
	levels         map[int]Level
	currentLevel   int
	debugLevel     uint8
	enemyBag       EnemyBag
	playerHand     Hand
	deck           Deck
	playerCastle   Castle
	shaders        map[ShaderName]*rl.Shader
	plainTileModel rl.Model
	waterTileModel rl.Model
	enemyModels    map[EnemyType]*rl.Model
	// cardModels     map[CardType]*rl.Model
	cardTextures map[CardType]*rl.Texture2D
	sunLight     Light
	spotLight    Light
	frameCount   int
	UI           UI
}

func Run() {
	game := &Game{}
	game.init()
	defer rl.CloseWindow()
	game.Loop()
}

func (g *Game) init() {
	g.CheckAndLoadConfig(true)
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagMsaa4xHint)
	rl.SetTargetFPS(g.Config.Window.TargetFPS)
	rl.InitWindow(g.Config.Window.Width, g.Config.Window.Height, g.Config.GameName)
	g.debugLevel = 0
	g.initCamera()
	g.levels = make(map[int]Level)
	g.tiles = make(map[TileType]Tile)
	g.shaders = make(map[ShaderName]*rl.Shader)
	g.enemyModels = make(map[EnemyType]*rl.Model)
	// g.cardModels = make(map[CardType]*rl.Model)
	g.cardTextures = make(map[CardType]*rl.Texture2D)
	g.LoadResources()
	g.initShadersAndLights()
	g.Round = g.NewRound()
	g.State = StatePlaying
}

func (g *Game) NextTurn() {
	g.UpdateEnemies()
}
