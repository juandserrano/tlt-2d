package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) Loop() {
	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		g.handleInput()
		g.player.update(dt)

		rl.BeginDrawing()
		rl.BeginMode2D(g.camera)
		rl.ClearBackground(rl.Gray)
		g.Draw()
		rl.EndMode2D()
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
