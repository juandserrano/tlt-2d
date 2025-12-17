package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StatePause
	StateGameOver
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
	camera         rl.Camera3D
	tiles          map[TileType]Tile
	levels         map[int]Level
	currentLevel   int
	debug          bool
	playerCastle   Castle
	ambientShader  rl.Shader
	waterShader    rl.Shader
	plainTileModel rl.Model
	waterTileModel rl.Model
	pawnModel      rl.Model
	sunLight       Light
	spotLight      Light
	testPawn       Enemy
	frameCount     int
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
	fmt.Println(g.Config.Window.TargetFPS)
	rl.InitWindow(g.Config.Window.Width, g.Config.Window.Height, g.Config.GameName)
	g.debug = false
	g.initCamera()
	g.levels = make(map[int]Level)
	g.tiles = make(map[TileType]Tile)
	g.LoadResources()

	g.State = StatePlaying
	g.Turn = TurnPlayer
	g.LoadLevelTiles(1)
	g.initPlayerCastle()
	g.NewEnemy(EnemyTypePawn, 5, 5)
	g.testPawn = EnemiesInPlay[0]

	g.initShadersAndLights()
}

func (g *Game) NextTurn() {
	g.UpdateEnemies()
}
