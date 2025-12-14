package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(int32(g.wWidth), int32(g.wHeight), "The Last Tower")
	g.camera = rl.NewCamera3D(rl.Vector3{X: 20, Y: 20, Z: 20}, rl.Vector3{X: 0, Y: 0, Z: 0}, rl.Vector3{0, 1, 0}, 70.0, rl.CameraPerspective)
	g.cameraMoveSpeed = CAMERA_MOVE_SPEED
	g.LoadBasicTile()
	g.levels = make(map[int]Level)
	g.currentLevel = 1
	// g.initPlayer()
	g.InitLevel1()

	// Load shader
	g.ambientShader = rl.LoadShader("assets/shaders/lighting2.vs", "assets/shaders/lighting2.fs")

	// Set shader location for standard attributes
	// Note: Raylib models usually have these bound by default.
	// In this simple case, Raylib handles standard attributes automatically when drawing models.

	// Ambient light level
	ambientLoc := rl.GetShaderLocation(g.ambientShader, "ambient")
	locViewPos := rl.GetShaderLocation(g.ambientShader, "viewPos")
	ambient := []float32{0.5, 0.5, 0.5, 1.0}
	rl.SetShaderValue(g.ambientShader, ambientLoc, ambient, rl.ShaderUniformVec4)
	rl.SetShaderValue(g.ambientShader, locViewPos, []float32{g.camera.Position.X, g.camera.Position.Y, g.camera.Position.Z}, rl.ShaderUniformVec3)
	g.sunLight = CreateLight(g.ambientShader, 0, LightDirectional, rl.NewVector3(0, 2, 0), rl.Vector3Zero(), rl.Blue, 20)
	CreateLight(g.ambientShader, 1, LightPoint, rl.NewVector3(1, 1, 1), rl.Vector3Zero(), rl.Green, 3)

	// Assign shader to all materials
	materials := g.basicTileModel.GetMaterials()
	for i := range materials {
		materials[i].Shader = g.ambientShader
	}
	fmt.Println("enabled:", g.sunLight.Enabled, " - color:", g.sunLight.Color, " - pos:", g.sunLight.Position, " - target:", g.sunLight.Target, " - type:", g.sunLight.Type)
}
