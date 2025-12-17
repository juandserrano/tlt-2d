package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) toggleDebug() {
	if rl.IsKeyPressed(rl.KeyG) {
		g.debugLevel++
		if g.debugLevel > 2 {
			g.debugLevel = 0
		}
	}
}
func (g *Game) DrawWorldDebug() {
	// Draw grid coords
	if g.debugLevel == 1 {
		for i := range g.levels[g.currentLevel].tiles {
			g.levels[g.currentLevel].tiles[i].debugDrawGridCoord(rl.Red)
		}
	}
}

func (g *Game) DrawStaticDebug() {
	rl.DrawFPS(10, 10)
	rl.DrawText("DEBUG MODE", int32(g.Config.Window.Width)-100, 10, 16, rl.Red)
	rl.DrawText(fmt.Sprintf("mousepos: %v", rl.GetMousePosition()), 100, 100, 20, rl.Blue)
}
