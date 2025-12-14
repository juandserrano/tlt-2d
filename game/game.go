package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WINDOW_WIDTH      = 800
	WINDOW_HEIGHT     = 600
	CAMERA_MOVE_SPEED = 3.0
)

type Game struct {
	wWidth          int
	wHeight         int
	camera          rl.Camera3D
	tiles           map[TileType]Tile
	levels          map[int]Level
	currentLevel    int
	cameraMoveSpeed float32
	debug           bool
	player          Player
	ambientShader   rl.Shader
	sunLight        Light
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
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagMsaa4xHint | rl.FlagVsyncHint)
	rl.SetTargetFPS(60)
	rl.InitWindow(int32(g.wWidth), int32(g.wHeight), "The Last Tower")
	g.initCamera()
	g.levels = make(map[int]Level)
	g.tiles = make(map[TileType]Tile)
	g.LoadResources()

	g.LoadLevel(1)
	g.initPlayer()

	g.initShadersAndLights()
}
