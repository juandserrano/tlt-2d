package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) handleInput() {
	g.handleCamera()
	g.handlePlayerMovement()
	g.toggleDebug()
}

func (g *Game) handlePlayerMovement() {

	if rl.IsKeyPressed(rl.KeyW) {
		g.MovePlayer(0, -1)
	}
	if rl.IsKeyPressed(rl.KeyUp) {
		g.MovePlayer(0, -1)
	}
	if rl.IsKeyPressed(rl.KeyS) {
		g.MovePlayer(0, 1)
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		g.MovePlayer(0, 1)
	}
	if rl.IsKeyPressed(rl.KeyA) {
		g.MovePlayer(-1, 0)
	}
	if rl.IsKeyPressed(rl.KeyLeft) {
		g.MovePlayer(-1, 0)
	}
	if rl.IsKeyPressed(rl.KeyD) {
		g.MovePlayer(1, 0)
	}
	if rl.IsKeyPressed(rl.KeyRight) {
		g.MovePlayer(1, 0)
	}
}
