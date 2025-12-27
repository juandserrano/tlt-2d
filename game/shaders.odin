package game

import "core:math"
import rl "vendor:raylib"

ShaderName :: enum {
	AmbientShader,
	WaterShader,
}

outlineShader: rl.Shader
initShadersAndLights :: proc(g: ^Game) {
	// SHADERS

	// Load ambient + diffuse shader
	g.shaders[.AmbientShader] = rl.LoadShader(
		"assets/shaders/lighting.vs",
		"assets/shaders/lighting.fs",
	)

	// Load outline shader
	outlineShader = rl.LoadShader("", "assets/shaders/glow.fs")
	// outlineShader = rl.LoadShader("", "assets/shaders/outline.fs")
	// Set Uniforms (Do this once if the texture size never changes)
	locSize := rl.GetShaderLocation(outlineShader, "textureSize")
	locRadius := rl.GetShaderLocation(outlineShader, "radius")
	locColor := rl.GetShaderLocation(outlineShader, "glowColor")
	rl.SetShaderValue(outlineShader, locRadius, &[]f32{9.0}, .FLOAT)
	rl.SetShaderValue(
		outlineShader,
		locSize,
		&[]f32 {
			f32(g.cardTextures[.CardTypeAttackPawn].width),
			f32(g.cardTextures[.CardTypeAttackPawn].height),
		},
		.VEC2,
	)
	// Set Outline Color (Red)
	rl.SetShaderValue(outlineShader, locColor, &[]f32{0.5, 0.0, 0.0}, .VEC3)

	// Ambient light level
	ambientLoc := rl.GetShaderLocation(g.shaders[.AmbientShader], "ambient")
	ambient := []f32{0.5, 0.5, 0.5, 1.0}
	rl.SetShaderValue(g.shaders[.AmbientShader], ambientLoc, &ambient, .VEC4)


	// g.playerCastle.model.materials.shader = g.shaders[.AmbientShader] // Assigh ambient shader to player tower
	// mats := g.playerCastle.model.GetMaterials()
	// for i := range mats {
	// 	mats[i].Shader = *g.shaders[AmbientShader]
	// }

	// Assign ambient shader to pawn model
	// mats = g.enemyModels[EnemyTypePawn].GetMaterials()
	// for j := range mats {
	// 	mats[j].Shader = *g.shaders[AmbientShader]
	// }

	// Assign ambient shader to knight model
	// mats = g.enemyModels[EnemyTypeKnight].GetMaterials()
	// for j := range mats {
	// 	mats[j].Shader = *g.shaders[AmbientShader]
	// }

	// Assign ambient shader to bishop model
	// mats = g.enemyModels[EnemyTypeBishop].GetMaterials()
	// for j := range mats {
	// 	mats[j].Shader = *g.shaders[AmbientShader]
	// }

	// Assign ambient shader to queen model
	// mats = g.enemyModels[EnemyTypeQueen].GetMaterials()
	// for j := range mats {
	// 	mats[j].Shader = *g.shaders[AmbientShader]
	// }

	// Assign ambient shader to king model
	// mats = g.enemyModels[EnemyTypeKing].GetMaterials()
	// for j := range mats {
	// 	mats[j].Shader = *g.shaders[AmbientShader]
	// }
	// Assign ambient shader to cards models
	// mats = g.cardModels[CardTypeAttackPawn].GetMaterials()
	// for i := range mats {
	// 	mats[i].Shader = *g.shaders[AmbientShader]
	// }

	g.shaders[.WaterShader] = rl.LoadShader("assets/shaders/water.vs", "assets/shaders/water.fs")

	// materials := g.tiles[TileTypeWater].model.GetMaterials()
	// for i := range materials {
	// 	materials[i].Shader = *g.shaders[WaterShader]
	// }

	// LIGTHS

	// Create basic sun illumination
	g.sunLight = CreateLight(
		g.shaders[.AmbientShader],
		0,
		.LightDirectional,
		rl.Vector3 {
			g.levels[g.currentLevel].centerXZ.x - 100,
			50,
			g.levels[g.currentLevel].centerXZ.y + 2,
		},
		rl.Vector3(0),
		rl.WHITE,
		2,
	)

	g.spotLight = CreateLight(
		g.shaders[.AmbientShader],
		1,
		.LightPoint,
		rl.Vector3{g.playerCastle.position.x - 3, 10, g.playerCastle.position.z},
		rl.Vector3{0, -1, 0},
		rl.WHITE,
		1,
	)

}

UpdateShaders :: proc(g: ^Game) {
	time := f32(rl.GetTime())

	// Animate sun (Circle around center)
	if g.Config.World.AnimateSun {
		AnimateSun(g, time)
	}

	// Get shader locations
	timeLoc := rl.GetShaderLocation(g.shaders[.WaterShader], "time")
	viewPosLoc := rl.GetShaderLocation(g.shaders[.WaterShader], "viewPos")
	rl.SetShaderValue(g.shaders[.WaterShader], timeLoc, &[]f32{time}, .FLOAT)

	camPos := []f32{g.camera.position.x, g.camera.position.y, g.camera.position.z}
	rl.SetShaderValue(g.shaders[.WaterShader], viewPosLoc, &camPos, .VEC3)

	// Update Outline Shader Time
	outlineTimeLoc := rl.GetShaderLocation(outlineShader, "time")
	rl.SetShaderValue(outlineShader, outlineTimeLoc, &[]f32{time}, .FLOAT)
}

AnimateSun :: proc(g: ^Game, time: f32) {
	g.sunLight.Position.x = f32(math.cos(f64(time)) * 100.0)
	g.sunLight.Position.z = f32(math.sin(f64(time)) * 50.0)
	UpdateLightValues(g.shaders[.AmbientShader], g.sunLight)
}
