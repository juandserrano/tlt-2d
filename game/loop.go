package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) Loop() {
	for !rl.WindowShouldClose() {
		g.handleCamera()
		rl.BeginDrawing()
		rl.BeginMode2D(g.camera)
		rl.ClearBackground(rl.Gray)
		g.Draw()
		rl.EndMode2D()
		rl.EndDrawing()
	}
}

func (g *Game) handleCamera() {

	wheel := rl.GetMouseWheelMove()
	if wheel != 0 {
		g.camera.Zoom += wheel * 0.1
		if g.camera.Zoom < 0.1 {
			g.camera.Zoom = 0.1
		}
	}

	moveSpeed := g.cameraMoveSpeed / g.camera.Zoom
	if rl.IsKeyDown(rl.KeyW) {
		g.camera.Target.Y -= moveSpeed
	}
	if rl.IsKeyDown(rl.KeyUp) {
		g.camera.Target.Y -= moveSpeed
	}
	if rl.IsKeyDown(rl.KeyS) {
		g.camera.Target.Y += moveSpeed
	}
	if rl.IsKeyDown(rl.KeyDown) {
		g.camera.Target.Y += moveSpeed
	}
	if rl.IsKeyDown(rl.KeyA) {
		g.camera.Target.X -= moveSpeed
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		g.camera.Target.X -= moveSpeed
	}
	if rl.IsKeyDown(rl.KeyD) {
		g.camera.Target.X += moveSpeed
	}
	if rl.IsKeyDown(rl.KeyRight) {
		g.camera.Target.X += moveSpeed
	}

}

func (g *Game) Draw() {
	g.DrawLevel(g.currentLevel)
}

func (g *Game) DrawLevel(currentLevel int) {
	level := g.levels[currentLevel]
	level.Draw()
}
