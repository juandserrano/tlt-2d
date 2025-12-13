package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) handleCamera() {
	// Zoom
	wheel := rl.GetMouseWheelMove()
	if wheel != 0 {
		// Raylib's UpdateCameraPro handles zoom via the 3rd argument (zoom factor)
		// Positive zooms in, negative zooms out.
		rl.UpdateCameraPro(&g.camera, rl.Vector3Zero(), rl.Vector3Zero(), -wheel)
	}

	// Orbit
	if rl.IsMouseButtonDown(rl.MouseMiddleButton) {
		delta := rl.GetMouseDelta()
		rotationSensitivity := float32(0.2)
		// UpdateCameraPro(camera, movement, rotation, zoom)
		// Rotation X: Horizontal (around UP axis)
		// Rotation Y: Vertical (around RIGHT axis)
		rl.UpdateCameraPro(&g.camera,
			rl.Vector3Zero(),
			rl.Vector3{X: delta.X * rotationSensitivity, Y: delta.Y * rotationSensitivity, Z: 0},
			0.0)
	}

	// Focus on player (optional, keep target on player?)
	// If we are orbiting freely, maybe we don't want to snap target to player every frame?
	// But usually "Orbit" implies orbiting THE PLAYER.
	// If we blindly update Target, it might fight UpdateCameraPro which modifies Position AND Target (relative).
	// Actually UpdateCameraPro keeps distance consistent?
	// Let's ensure Target closely follows player if desired, or we just let it be free.
	// User didn't specify "Follow Player", just "User can orbit".
	// But previous code had `moveCameraWithPlayer`.
	// For now, let's leave it loose. If the user wants to follow player, they can ask.
	// Or simplistic: Update Target to player position, but recalculate Position to maintain offset?
	// That's complex. Let's just allow free orbit around current target.

	// Panning with Right Click? (Optional, good for UX)
	if rl.IsMouseButtonDown(rl.MouseRightButton) {
		delta := rl.GetMouseDelta()
		panSpeed := float32(0.05)
		// Panning moves both Position and Target
		// UpdateCameraPro input movement is relative to camera frame?
		// Forward/Back, Right/Left, Up/Down.
		// Delta X -> Right/Left
		// Delta Y -> Up/Down (or Forward/Back depending on view)
		// Let's implement simple pan:
		// Right = -delta.X, Up = delta.Y
		rl.UpdateCameraPro(&g.camera,
			rl.Vector3{X: -delta.X * panSpeed, Y: delta.Y * panSpeed, Z: 0},
			rl.Vector3Zero(),
			0.0)
	}
}

func moveCameraWithPlayer(g *Game) {
	g.camera.Target = g.player.position
}

func moveCameraWithLimits(g *Game) {
	moveSpeed := g.cameraMoveSpeed / g.camera.Fovy
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
		if t.x == 0 && t.z == 39 {
			leftMostPos = t.position.X
		}
		if t.x == 39 && t.z == 0 {
			rightMostPos = t.position.X
		}
		if t.x == 39 && t.z == 39 {
			bottomMostPos = t.position.Y
		}
	}
	if g.camera.Target.X < leftMostPos {
		g.camera.Target.X = leftMostPos
	}
	if g.camera.Target.X > rightMostPos-float32(rl.GetScreenWidth())*g.camera.Fovy {
		g.camera.Target.X = rightMostPos - float32(rl.GetScreenWidth())*g.camera.Fovy
	}
	if g.camera.Target.Y < -100 {
		g.camera.Target.Y = -100
	}
	if g.camera.Target.Y > bottomMostPos-float32(rl.GetScreenHeight()) {
		g.camera.Target.Y = bottomMostPos - float32(rl.GetScreenHeight())
	}

}
