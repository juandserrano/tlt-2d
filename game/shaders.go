package game

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ShaderName int

const (
	AmbientShader ShaderName = iota
	WaterShader
)

var outlineShader rl.Shader

func (g *Game) initShadersAndLights() {
	// SHADERS

	// Load ambient + diffuse shader
	aShader := new(rl.Shader)
	*aShader = rl.LoadShader("assets/shaders/lighting.vs", "assets/shaders/lighting.fs")
	g.shaders[AmbientShader] = aShader

	// Load outline shader
	outlineShader = rl.LoadShader("", "assets/shaders/glow.fs")
	// outlineShader = rl.LoadShader("", "assets/shaders/outline.fs")
	// Set Uniforms (Do this once if the texture size never changes)
	locSize := rl.GetShaderLocation(outlineShader, "textureSize")
	locRadius := rl.GetShaderLocation(outlineShader, "radius")
	locColor := rl.GetShaderLocation(outlineShader, "glowColor")
	rl.SetShaderValue(outlineShader, locRadius, []float32{9.0}, rl.ShaderUniformFloat)
	rl.SetShaderValue(outlineShader, locSize, []float32{float32(g.cardTextures[CardTypeAttackPawn].Width), float32(g.cardTextures[CardTypeAttackPawn].Height)}, rl.ShaderUniformVec2)
	// Set Outline Color (Red)
	rl.SetShaderValue(outlineShader, locColor, []float32{0.5, 0.0, 0.0}, rl.ShaderUniformVec3)

	// Ambient light level
	ambientLoc := rl.GetShaderLocation(*g.shaders[AmbientShader], "ambient")
	ambient := []float32{0.5, 0.5, 0.5, 1.0}
	rl.SetShaderValue(*g.shaders[AmbientShader], ambientLoc, ambient, rl.ShaderUniformVec4)

	// Assigh ambient shader to player tower
	mats := g.playerCastle.model.GetMaterials()
	for i := range mats {
		mats[i].Shader = *g.shaders[AmbientShader]
	}

	// Assign ambient shader to pawn model
	mats = g.enemyModels[EnemyTypePawn].GetMaterials()
	for j := range mats {
		mats[j].Shader = *g.shaders[AmbientShader]
	}

	// Assign ambient shader to knight model
	mats = g.enemyModels[EnemyTypeKnight].GetMaterials()
	for j := range mats {
		mats[j].Shader = *g.shaders[AmbientShader]
	}

	// Assign ambient shader to bishop model
	mats = g.enemyModels[EnemyTypeBishop].GetMaterials()
	for j := range mats {
		mats[j].Shader = *g.shaders[AmbientShader]
	}

	// Assign ambient shader to queen model
	mats = g.enemyModels[EnemyTypeQueen].GetMaterials()
	for j := range mats {
		mats[j].Shader = *g.shaders[AmbientShader]
	}

	// Assign ambient shader to king model
	mats = g.enemyModels[EnemyTypeKing].GetMaterials()
	for j := range mats {
		mats[j].Shader = *g.shaders[AmbientShader]
	}
	// Assign ambient shader to cards models
	// mats = g.cardModels[CardTypeAttackPawn].GetMaterials()
	// for i := range mats {
	// 	mats[i].Shader = *g.shaders[AmbientShader]
	// }

	wShader := new(rl.Shader)
	*wShader = rl.LoadShader("assets/shaders/water.vs", "assets/shaders/water.fs")
	g.shaders[WaterShader] = wShader

	materials := g.tiles[TileTypeWater].model.GetMaterials()
	for i := range materials {
		materials[i].Shader = *g.shaders[WaterShader]
	}

	// LIGTHS

	// Create basic sun illumination
	g.sunLight = CreateLight(
		*g.shaders[AmbientShader], 0, LightDirectional,
		rl.NewVector3(g.levels[g.currentLevel].centerXZ.X-100, 50, g.levels[g.currentLevel].centerXZ.Y+2),
		rl.Vector3Zero(),
		rl.White,
		2)

	g.spotLight = CreateLight(
		*g.shaders[AmbientShader], 1, LightPoint,
		rl.NewVector3(g.playerCastle.position.X-3, 10, g.playerCastle.position.Z),
		rl.Vector3{X: 0, Y: -1, Z: 0},
		rl.White,
		1)

}

func (g *Game) UpdateShaders() {
	time := float32(rl.GetTime())

	// Animate sun (Circle around center)
	if g.Config.World.AnimateSun {
		g.AnimateSun(time)
	}

	// Get shader locations
	timeLoc := rl.GetShaderLocation(*g.shaders[WaterShader], "time")
	viewPosLoc := rl.GetShaderLocation(*g.shaders[WaterShader], "viewPos")
	rl.SetShaderValue(*g.shaders[WaterShader], timeLoc, []float32{time}, rl.ShaderUniformFloat)

	camPos := []float32{g.camera.Position.X, g.camera.Position.Y, g.camera.Position.Z}
	rl.SetShaderValue(*g.shaders[WaterShader], viewPosLoc, camPos, rl.ShaderUniformVec3)

	// Update Outline Shader Time
	outlineTimeLoc := rl.GetShaderLocation(outlineShader, "time")
	rl.SetShaderValue(outlineShader, outlineTimeLoc, []float32{time}, rl.ShaderUniformFloat)
}

func (g *Game) AnimateSun(time float32) {
	g.sunLight.Position.X = float32(math.Cos(float64(time)) * 100.0)
	g.sunLight.Position.Z = float32(math.Sin(float64(time)) * 50.0)
	UpdateLightValues(*g.shaders[AmbientShader], g.sunLight)
}
