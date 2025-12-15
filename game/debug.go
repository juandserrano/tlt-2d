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
	rl.DrawSphere(rl.Vector3{5, 0, 5}, 0.1, rl.Red)
}

func (g *Game) DrawStaticDebug() {
	rl.DrawFPS(10, 10)
	rl.DrawText("DEBUG MODE", int32(g.wWidth)-100, 10, 16, rl.Red)
	rl.DrawText(fmt.Sprintf("%d, %d", EnemiesInPlay[0].gridPos.X, EnemiesInPlay[0].gridPos.Z), 10, 25, 20, rl.LightGray)
}
