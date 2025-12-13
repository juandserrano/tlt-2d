package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) Loop() {
	// rl.DisableCursor()
	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&g.camera, rl.CameraFree)
		dt := rl.GetFrameTime()
		g.handleInput(dt)
		g.player.update(dt)

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
