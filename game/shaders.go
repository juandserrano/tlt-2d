package game

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) initShadersAndLights() {
	// SHADERS

	// Load ambient + diffuse shader
	g.ambientShader = rl.LoadShader("assets/shaders/lighting.vs", "assets/shaders/lighting.fs")

	// Ambient light level
	ambientLoc := rl.GetShaderLocation(g.ambientShader, "ambient")
	ambient := []float32{0.5, 0.5, 0.5, 1.0}
	rl.SetShaderValue(g.ambientShader, ambientLoc, ambient, rl.ShaderUniformVec4)

	// Assigh ambient shader to player tower
	mats := g.playerCastle.model.GetMaterials()
	for i := range mats {
		mats[i].Shader = g.ambientShader
	}

	g.waterShader = rl.LoadShader("assets/shaders/water.vs", "assets/shaders/water.fs")

	materials := g.tiles[TileTypeWater].model.GetMaterials()
	for i := range materials {
		materials[i].Shader = g.waterShader
	}

	// LIGTHS

	// Create basic sun illumination
	g.sunLight = CreateLight(
		g.ambientShader, 0, LightDirectional,
		rl.NewVector3(g.levels[g.currentLevel].centerXZ.X-2, 50, g.levels[g.currentLevel].centerXZ.Y+2),
		rl.Vector3Zero(),
		rl.White,
		1)

	g.spotLight = CreateLight(
		g.ambientShader, 1, LightPoint,
		rl.NewVector3(g.playerCastle.position.X-3, 5, g.playerCastle.position.Z),
		rl.Vector3{X: 0, Y: -1, Z: 0},
		rl.White,
		2)

}

func (g *Game) UpdateShaders() {
	time := float32(rl.GetTime())

	// Animate sun (Circle around center)
	if g.Config.World.AnimateSun {
		g.AnimateSun(time)
	}

	// Get shader locations
	timeLoc := rl.GetShaderLocation(g.waterShader, "time")
	viewPosLoc := rl.GetShaderLocation(g.waterShader, "viewPos")
	rl.SetShaderValue(g.waterShader, timeLoc, []float32{time}, rl.ShaderUniformFloat)

	camPos := []float32{g.camera.Position.X, g.camera.Position.Y, g.camera.Position.Z}
	rl.SetShaderValue(g.waterShader, viewPosLoc, camPos, rl.ShaderUniformVec3)
}

func (g *Game) AnimateSun(time float32) {
	g.sunLight.Position.X = float32(math.Cos(float64(time)) * 10.0)
	g.sunLight.Position.Z = float32(math.Sin(float64(time)) * 5.0)
	UpdateLightValues(g.ambientShader, g.sunLight)
}
