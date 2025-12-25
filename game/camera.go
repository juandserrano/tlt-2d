package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) initCamera() {
	tar := rl.NewVector3(0, 0, 0)
	g.camera = rl.NewCamera3D(
		rl.Vector3{X: tar.X - 5, Y: 15, Z: tar.Z},
		tar,
		rl.Vector3{X: 0, Y: 1, Z: 0}, 70.0,
		rl.CameraPerspective)
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

	// Camera Shake
	if g.CameraShakeIntensity > 0 {
		g.CameraShakeIntensity -= 0.05 // Decay
		if g.CameraShakeIntensity < 0 {
			g.CameraShakeIntensity = 0
		}
	}
}

func (g *Game) GetRenderCamera() rl.Camera3D {
	cam := g.camera
	if g.CameraShakeIntensity > 0 {
		offset := rl.Vector3{
			X: (float32(rl.GetRandomValue(0, 100))/50.0 - 1.0) * g.CameraShakeIntensity,
			Y: (float32(rl.GetRandomValue(0, 100))/50.0 - 1.0) * g.CameraShakeIntensity,
			Z: (float32(rl.GetRandomValue(0, 100))/50.0 - 1.0) * g.CameraShakeIntensity,
		}
		cam.Position = rl.Vector3Add(cam.Position, offset)
		cam.Target = rl.Vector3Add(cam.Target, offset)
	}
	return cam
}
