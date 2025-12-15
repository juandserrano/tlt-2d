package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) initCamera() {
	tar := rl.NewVector3(0, 0, 0)
	g.camera = rl.NewCamera3D(
		rl.Vector3{X: tar.X - 5, Y: 10, Z: tar.Z},
		tar,
		rl.Vector3{X: 0, Y: 1, Z: 0}, 70.0,
		rl.CameraPerspective)
	g.cameraMoveSpeed = CAMERA_MOVE_SPEED
}

func (g *Game) handleCamera() {
	// Zoom
	wheel := rl.GetMouseWheelMove()
	if wheel != 0 {
		rl.UpdateCameraPro(&g.camera, rl.Vector3Zero(), rl.Vector3Zero(), -wheel)
	}

	// Orbit
	if rl.IsMouseButtonDown(rl.MouseMiddleButton) {
		delta := rl.GetMouseDelta()
		rotationSensitivity := float32(0.2)
		rl.UpdateCameraPro(&g.camera,
			rl.Vector3Zero(),
			rl.Vector3{X: delta.X * rotationSensitivity, Y: delta.Y * rotationSensitivity, Z: 0},
			0.0)
	}

	// Panning with Right Click? (Optional, good for UX)
	if rl.IsMouseButtonDown(rl.MouseRightButton) {
		delta := rl.GetMouseDelta()
		panSpeed := float32(0.05)
		rl.UpdateCameraPro(&g.camera,
			rl.Vector3{X: delta.Y * panSpeed, Y: -delta.X * panSpeed, Z: 0},
			rl.Vector3Zero(),
			0.0)
	}
}

// func moveCameraWithPlayer(g *Game) {
// 	g.camera.Target = g.player.position
// }

// func moveCameraWithLimits(g *Game) {
// 	moveSpeed := g.cameraMoveSpeed / g.camera.Fovy
// 	const sideBufferPx = 20
// 	mousePosition := rl.GetMousePosition()
// 	if mousePosition.X >= float32(rl.GetScreenWidth())-sideBufferPx {
// 		g.camera.Target.X += moveSpeed
// 	}
// 	if mousePosition.X <= sideBufferPx {
// 		g.camera.Target.X -= moveSpeed
// 	}
// 	if mousePosition.Y >= float32(rl.GetScreenHeight())-sideBufferPx {
// 		g.camera.Target.Y += moveSpeed
// 	}
// 	if mousePosition.Y <= sideBufferPx {
// 		g.camera.Target.Y -= moveSpeed
// 	}

// 	var leftMostPos float32
// 	var rightMostPos float32
// 	var bottomMostPos float32
// 	for _, t := range g.levels[g.currentLevel].tiles {
// 		if t.x == 0 && t.z == 39 {
// 			leftMostPos = t.position.X
// 		}
// 		if t.x == 39 && t.z == 0 {
// 			rightMostPos = t.position.X
// 		}
// 		if t.x == 39 && t.z == 39 {
// 			bottomMostPos = t.position.Y
// 		}
// 	}
// 	if g.camera.Target.X < leftMostPos {
// 		g.camera.Target.X = leftMostPos
// 	}
// 	if g.camera.Target.X > rightMostPos-float32(rl.GetScreenWidth())*g.camera.Fovy {
// 		g.camera.Target.X = rightMostPos - float32(rl.GetScreenWidth())*g.camera.Fovy
// 	}
// 	if g.camera.Target.Y < -100 {
// 		g.camera.Target.Y = -100
// 	}
// 	if g.camera.Target.Y > bottomMostPos-float32(rl.GetScreenHeight()) {
// 		g.camera.Target.Y = bottomMostPos - float32(rl.GetScreenHeight())
// 	}

// }
