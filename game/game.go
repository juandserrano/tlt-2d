package game

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	WINDOW_WIDTH      = 1024
	WINDOW_HEIGHT     = 720
	CAMERA_MOVE_SPEED = 3.0
)

type Game struct {
	wWidth  int
	wHeight int
	// camera          rl.Camera2D
	camera rl.Camera3D
	// tex             rl.Texture2D
	basicTileModel  rl.Model
	levels          map[int]Level
	currentLevel    int
	cameraMoveSpeed float32
	debug           bool
	player          Player
}

func Run() {
	game := &Game{}
	game.init()
	defer rl.CloseWindow()
	game.Loop()

}

func (g *Game) init() {
	g.wWidth = WINDOW_WIDTH
	g.wHeight = WINDOW_HEIGHT
	g.debug = false
	rl.SetConfigFlags(rl.FlagWindowResizable)
	// rl.ClearWindowState(rl.FlagWindowResizable)
	rl.InitWindow(int32(g.wWidth), int32(g.wHeight), "The Last Tower")
	g.camera = rl.NewCamera3D(rl.Vector3{X: 20, Y: 20, Z: 20}, rl.Vector3{X: 0, Y: 0, Z: 0}, rl.Vector3{0, 1, 0}, 70.0, rl.CameraPerspective)
	g.cameraMoveSpeed = CAMERA_MOVE_SPEED
	// g.tex = rl.LoadTexture("assets/grass_tile.png")
	g.LoadBasicTile()
	g.levels = make(map[int]Level)
	g.currentLevel = 1
	// g.initPlayer()
	g.InitLevel1()
}
