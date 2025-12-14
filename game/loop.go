package game

import (
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

		rl.BeginDrawing()
		rl.ClearBackground(rl.Gray)
		rl.BeginMode3D(g.camera)
		rl.DrawGrid(10, 1.0)
		g.Draw()
		rl.EndMode3D()
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
