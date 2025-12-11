package game

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	WINDOW_WIDTH      = 1024
	WINDOW_HEIGHT     = 720
	CAMERA_MOVE_SPEED = 3.0
)

type Game struct {
	wWidth          int
	wHeight         int
	camera          rl.Camera2D
	tex             rl.Texture2D
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
	// rl.ClearWindowState(rl.FlagWindowResizable)
	rl.InitWindow(int32(g.wWidth), int32(g.wHeight), "The Last Tower")
	g.camera = rl.NewCamera2D(rl.Vector2{X: WINDOW_WIDTH / 2, Y: WINDOW_HEIGHT / 2}, rl.Vector2{X: 0, Y: 0}, 0, 1)
	g.cameraMoveSpeed = CAMERA_MOVE_SPEED
	g.tex = rl.LoadTexture("assets/grass_tile.png")
	g.levels = make(map[int]Level)
	g.currentLevel = 1
	g.initPlayer()
	g.InitLevel1()
}
