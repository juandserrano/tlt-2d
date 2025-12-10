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
	DrawTileCenters(g)
}

func (g *Game) DrawStaticDebug() {
	zoomStr := fmt.Sprintf("Zoom: %.2f", g.camera.Zoom)
	rl.DrawText(zoomStr, 10, 10, 20, rl.Red)
}

func DrawTileCenters(g *Game) {
	for _, t := range g.levels[g.currentLevel].tiles {
		rl.DrawCircle(int32(t.position.X), int32(t.position.Y), 2, rl.Red)
	}
}
