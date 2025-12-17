package game

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Loop() {
	// rl.DisableCursor()
	frameCount := 0
	for !rl.WindowShouldClose() {
		frameCount++
		if frameCount%30 == 0 {
			g.CheckAndLoadConfig(false)
		}
		dt := rl.GetFrameTime()
		g.handleInput(dt)
		g.playerCastle.update(dt)

		time := float32(rl.GetTime())

		// Animate sun (Circle around center)
		// g.sunLight.Position.X = float32(math.Cos(float64(time)) * 10.0)
		// g.sunLight.Position.Z = float32(math.Sin(float64(time)) * 5.0)
		// UpdateLightValues(g.ambientShader, g.sunLight)

		// Get shader locations
		timeLoc := rl.GetShaderLocation(g.waterShader, "time")
		viewPosLoc := rl.GetShaderLocation(g.waterShader, "viewPos")
		rl.SetShaderValue(g.waterShader, timeLoc, []float32{time}, rl.ShaderUniformFloat)

		camPos := []float32{g.camera.Position.X, g.camera.Position.Y, g.camera.Position.Z}
		rl.SetShaderValue(g.waterShader, viewPosLoc, camPos, rl.ShaderUniformVec3)

		rl.BeginDrawing()
		rl.ClearBackground(color.RGBA{
			uint8(g.Config.Window.BackgroundColor.R * 255),
			uint8(g.Config.Window.BackgroundColor.G * 255),
			uint8(g.Config.Window.BackgroundColor.B * 255),
			uint8(g.Config.Window.BackgroundColor.A * 255)})
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
	g.playerCastle.draw()
	g.drawEnemies()
	if g.debug {
		g.DrawWorldDebug()
	}
}

func (g *Game) DrawLevel(currentLevel int) {
	level := g.levels[currentLevel]
	level.Draw()
}
