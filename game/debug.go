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
	// rl.DrawRectangleRec(rl.NewRectangle(g.player.position.X, g.player.position.Y, float32(g.player.texture.Width), float32(g.player.texture.Height)), rl.Blue)
}

func (g *Game) DrawStaticDebug() {
	zoomStr := fmt.Sprintf("Zoom: %.2f", g.camera.Zoom)
	rl.DrawText(zoomStr, 10, 10, 20, rl.Red)
	// playerGridPosStr := fmt.Sprintf("x: %d, y: %d", g.player.gridX, g.player.gridY)
	// rl.DrawText(playerGridPosStr, 10, 20, 20, rl.Red)
	playerPosStr := fmt.Sprintf("x: %.2f, y: %.2f", g.player.position.X, g.player.position.Y)
	rl.DrawText(playerPosStr, 10, 40, 20, rl.Red)
	delta := fmt.Sprintf("deltaX: %.2f, deltaY: %.2f", g.player.position.X-g.player.prevPos.X, g.player.position.Y-g.player.prevPos.Y)
	rl.DrawText(delta, 10, 80, 20, rl.Red)
}

func DrawTileCenters(g *Game) {
	for _, t := range g.levels[g.currentLevel].tiles {
		rl.DrawCircle(int32(t.position.X), int32(t.position.Y), 2, rl.Red)
		// coords := fmt.Sprintf("x: %.2f, y: %.2f", t.position.X, t.position.Y)
		// gridcoords := fmt.Sprintf("gx: %d, gy: %d", t.x, t.y)
		// rl.DrawText(coords, int32(t.position.X)+10, int32(t.position.Y), 4, rl.White)
		// rl.DrawText(gridcoords, int32(t.position.X)+10, int32(t.position.Y-10), 4, rl.White)
	}
}
