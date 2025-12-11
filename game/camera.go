package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) handleCamera() {
	// zoom
	wheel := rl.GetMouseWheelMove()
	if wheel != 0 {
		g.camera.Zoom += wheel * 0.1
		if g.camera.Zoom < 0.5 {
			g.camera.Zoom = 0.5
		}
		if g.camera.Zoom > 2 {
			g.camera.Zoom = 2
		}
	}

	focusOnPlayer := true
	if focusOnPlayer {
		moveCameraWithPlayer(g)
	} else {
		moveCameraWithLimits(g)
	}
}

func moveCameraWithPlayer(g *Game) {
	g.camera.Target = g.player.position
}

func moveCameraWithLimits(g *Game) {
	moveSpeed := g.cameraMoveSpeed / g.camera.Zoom
	const sideBufferPx = 20
	mousePosition := rl.GetMousePosition()
	if mousePosition.X >= float32(rl.GetScreenWidth())-sideBufferPx {
		g.camera.Target.X += moveSpeed
	}
	if mousePosition.X <= sideBufferPx {
		g.camera.Target.X -= moveSpeed
	}
	if mousePosition.Y >= float32(rl.GetScreenHeight())-sideBufferPx {
		g.camera.Target.Y += moveSpeed
	}
	if mousePosition.Y <= sideBufferPx {
		g.camera.Target.Y -= moveSpeed
	}

	var leftMostPos float32
	var rightMostPos float32
	var bottomMostPos float32
	for _, t := range g.levels[g.currentLevel].tiles {
		if t.x == 0 && t.y == 39 {
			leftMostPos = t.position.X
		}
		if t.x == 39 && t.y == 0 {
			rightMostPos = t.position.X
		}
		if t.x == 39 && t.y == 39 {
			bottomMostPos = t.position.Y
		}
	}
	if g.camera.Target.X < leftMostPos {
		g.camera.Target.X = leftMostPos
	}
	if g.camera.Target.X > rightMostPos-float32(rl.GetScreenWidth())*g.camera.Zoom {
		g.camera.Target.X = rightMostPos - float32(rl.GetScreenWidth())*g.camera.Zoom
	}
	if g.camera.Target.Y < -100 {
		g.camera.Target.Y = -100
	}
	if g.camera.Target.Y > bottomMostPos-float32(rl.GetScreenHeight()) {
		g.camera.Target.Y = bottomMostPos - float32(rl.GetScreenHeight())
	}

}
