package game

import (
	"unsafe"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Loop() {
	// rl.DisableCursor()
	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		g.handleInput(dt)
		g.player.update(dt)

		// ----------------------------------------------------------------------------------
		// Shadow Pass
		// ----------------------------------------------------------------------------------
		lightDir := rl.Vector3{X: -100, Y: 40, Z: -100}
		lightPos := lightDir // Directional light, we treat pos as source for matrix

		// Calculate Light View-Projection Matrix
		// View: Look from lightPos to 0,0,0
		matLightView := rl.MatrixLookAt(lightPos, rl.Vector3Zero(), rl.Vector3{X: 0, Y: 1, Z: 0})
		// Proj: Orthographic
		orthoSize := float32(40.0) // Increase coverage
		matLightProj := rl.MatrixOrtho(-orthoSize, orthoSize, -orthoSize, orthoSize, 0.1, 200.0)
		g.matLight = rl.MatrixMultiply(matLightView, matLightProj)

		// Render to Shadow Map
		rl.BeginTextureMode(g.shadowMap)
		rl.ClearBackground(rl.White)

		// Swap shader to shadow shader for all materials
		materials := unsafe.Slice(g.basicTileModel.Materials, g.basicTileModel.MaterialCount)
		for i := range materials {
			materials[i].Shader = g.shadowShader
		}

		// Draw scene from light perspective
		lightCamera := rl.Camera3D{
			Position:   lightPos,
			Target:     rl.Vector3Zero(),
			Up:         rl.Vector3{X: 0, Y: 1, Z: 0},
			Fovy:       orthoSize * 2,
			Projection: rl.CameraOrthographic,
		}

		rl.BeginMode3D(lightCamera)
		g.Draw()
		rl.EndMode3D()

		// Restore lighting shader
		for i := range materials {
			materials[i].Shader = g.shader
		}

		rl.EndTextureMode()

		// ----------------------------------------------------------------------------------
		// Main Pass Updates
		// ----------------------------------------------------------------------------------

		// Update uniforms
		viewPos := []float32{g.camera.Position.X, g.camera.Position.Y, g.camera.Position.Z}
		rl.SetShaderValue(g.shader, g.viewPosLoc, viewPos, rl.ShaderUniformVec3)

		rl.SetShaderValue(g.shader, g.lightPosLoc, []float32{lightDir.X, lightDir.Y, lightDir.Z}, rl.ShaderUniformVec3)

		// Send Light Matrix to shader
		matLightLoc := rl.GetShaderLocation(g.shader, "matLight")
		rl.SetShaderValueMatrix(g.shader, matLightLoc, g.matLight)

		// Bind Shadow Map to Material Map Slot 6 (Height)
		// This ensures Raylib binds it to Texture Unit 6 when drawing the model
		for i := range materials {
			// Access Maps array (assuming size > 6)
			// unsafe.Slice requires strict size knowledge or we risk crash if Materials.Maps is small?
			// Raylib-go allocates C-struct.
			// Default Material has enough maps.
			// We trust Maps points to valid array.
			maps := unsafe.Slice(materials[i].Maps, 11)
			maps[rl.MapHeight].Texture = g.shadowMap.Texture // Using Color buffer. If fails, check Depth accessibility.
			// maps[rl.MapHeight].Value = ... if needed? Default is suitable.
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Gray)
		rl.BeginMode3D(g.camera)
		// rl.BeginMode2D(g.camera)
		rl.DrawGrid(10, 1.0)
		g.Draw()
		rl.EndMode3D()
		// rl.EndMode2D()
		if g.debug {
			g.DrawStaticDebug()
		}
		rl.EndDrawing()
	}
}

func (g *Game) Draw() {
	g.DrawLevel(g.currentLevel)
	g.player.draw()
	if g.debug {
		g.DrawWorldDebug()
	}
}

func (g *Game) DrawLevel(currentLevel int) {
	level := g.levels[currentLevel]
	level.Draw()
}
