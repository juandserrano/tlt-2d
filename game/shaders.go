package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) initShadersAndLights() {
	// SHADERS

	// Load ambient + diffuse shader
	g.ambientShader = rl.LoadShader("assets/shaders/lighting.vs", "assets/shaders/lighting.fs")

	// Ambient light level
	ambientLoc := rl.GetShaderLocation(g.ambientShader, "ambient")
	ambient := []float32{0.5, 0.5, 0.5, 1.0}
	rl.SetShaderValue(g.ambientShader, ambientLoc, ambient, rl.ShaderUniformVec4)

	// Assign shader to all materials
	materials := g.basicTileModel.GetMaterials()
	for i := range materials {
		materials[i].Shader = g.ambientShader
	}

	// LIGTHS

	// Create basic sun illumination
	g.sunLight = CreateLight(
		g.ambientShader, 0, LightDirectional,
		rl.NewVector3(g.levels[g.currentLevel].centerXZ.X, 5, g.levels[g.currentLevel].centerXZ.Y),
		rl.Vector3Zero(),
		rl.White,
		2)

}
