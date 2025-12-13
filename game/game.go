package game

import (
	"unsafe"

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
	shader          rl.Shader
	viewPosLoc      int32
	lightPosLoc     int32
	shadowShader    rl.Shader
	shadowMap       rl.RenderTexture2D
	matLight        rl.Matrix
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

	// Load shader
	g.shader = rl.LoadShader("assets/shaders/lighting.vs", "assets/shaders/lighting.fs")
	g.viewPosLoc = rl.GetShaderLocation(g.shader, "viewPos")
	g.lightPosLoc = rl.GetShaderLocation(g.shader, "lightPos")

	// Update lighting shader with emission map location (texture unit 5, standard for MATERIAL_MAP_EMISSION)
	// Raylib LoadModel loads GLB materials, mapping Emission texture to index 5 if present.
	// When DrawModel is called, it binds Map[5].Texture to Texture Unit 5.
	// We tell the shader that 'texture_emission' sampler should read from Unit 5.
	emissionMapLoc := rl.GetShaderLocation(g.shader, "texture_emission")
	rl.SetShaderValue(g.shader, emissionMapLoc, []float32{rl.MapEmission}, rl.ShaderUniformInt)

	// Set shader location for standard attributes
	// Note: Raylib models usually have these bound by default.
	// In this simple case, Raylib handles standard attributes automatically when drawing models.

	// Ambient light level
	ambientLoc := rl.GetShaderLocation(g.shader, "ambient")
	ambient := []float32{0.4, 0.4, 0.4, 1.0}
	rl.SetShaderValue(g.shader, ambientLoc, ambient, rl.ShaderUniformVec4)

	// Assign shader to all materials
	materials := unsafe.Slice(g.basicTileModel.Materials, g.basicTileModel.MaterialCount)
	for i := range materials {
		materials[i].Shader = g.shader
	}

	g.shadowShader = rl.LoadShader("assets/shaders/shadow.vs", "assets/shaders/shadow.fs")

	g.shadowMap = rl.LoadRenderTexture(2048, 2048)

	// Update lighting shader with shadow map location (texture unit 6, using MapHeight slot)
	shadowMapLoc := rl.GetShaderLocation(g.shader, "shadowMap")
	rl.SetShaderValue(g.shader, shadowMapLoc, []float32{6}, rl.ShaderUniformInt)
}
