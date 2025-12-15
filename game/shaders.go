package game

import (
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

	// Assign ambient shader to all tile models
	for _, v := range g.tiles {
		materials := v.model.GetMaterials()
		for i := range materials {
			materials[i].Shader = g.ambientShader
		}
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

}
