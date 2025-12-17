package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) toggleDebug() {
	if rl.IsKeyPressed(rl.KeyG) {
		g.debug = !g.debug
	}
}
func (g *Game) DrawWorldDebug() {
	// Draw grid coords
	for i := range g.levels[g.currentLevel].tiles {
		g.levels[g.currentLevel].tiles[i].debugDrawGridCoord()
	}
}

func (g *Game) DrawStaticDebug() {
	rl.DrawFPS(10, 10)
	rl.DrawText("DEBUG MODE", int32(g.Config.Window.Width)-100, 10, 16, rl.Red)
	rl.DrawText(fmt.Sprintf("%d, %d", EnemiesInPlay[0].gridPos.X, EnemiesInPlay[0].gridPos.Z), 10, 25, 20, rl.LightGray)
}
