package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Loop() {
	// rl.DisableCursor()
	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&g.camera, rl.CameraFree)
		dt := rl.GetFrameTime()
		g.handleInput(dt)
		g.player.update(dt)

		// Update shader values
		viewPos := []float32{g.camera.Position.X, g.camera.Position.Y, g.camera.Position.Z}
		rl.SetShaderValue(g.shader, g.viewPosLoc, viewPos, rl.ShaderUniformVec3)

		// Fixed sun position
		// time := float64(rl.GetTime())
		lightPos := []float32{-10, 10, -10}
		rl.SetShaderValue(g.shader, g.lightPosLoc, lightPos, rl.ShaderUniformVec3)

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
