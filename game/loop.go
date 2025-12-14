package game

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Loop() {
	// rl.DisableCursor()
	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		g.handleInput(dt)
		g.player.update(dt)
		// Animate sun (Circle around center)
		time := float32(rl.GetTime())
		g.sunLight.Position.X = float32(math.Cos(float64(time)) * 10.0)
		g.sunLight.Position.Z = float32(math.Sin(float64(time)) * 5.0)
		UpdateLightValues(g.ambientShader, g.sunLight)

		// Get shader locations
		timeLoc := rl.GetShaderLocation(g.waterShader, "time")
		viewPosLoc := rl.GetShaderLocation(g.waterShader, "viewPos")
		rl.SetShaderValue(g.waterShader, timeLoc, []float32{time}, rl.ShaderUniformFloat)

		camPos := []float32{g.camera.Position.X, g.camera.Position.Y, g.camera.Position.Z}
		rl.SetShaderValue(g.waterShader, viewPosLoc, camPos, rl.ShaderUniformVec3)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Gray)
		rl.BeginMode3D(g.camera)
		rl.DrawGrid(10, 1.0)
		g.Draw()
		rl.EndMode3D()
		if g.debug {
			g.DrawStaticDebug()
			rl.DrawText(fmt.Sprintf("mousepos: %v", rl.GetMousePosition()), 100, 100, 20, rl.Blue)
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
