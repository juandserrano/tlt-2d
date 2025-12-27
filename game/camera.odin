package game

import rl "vendor:raylib"

initCamera :: proc(g: ^Game) {
	tar := rl.Vector3{0, 0, 0}
	g.camera = rl.Camera3D {
		rl.Vector3{tar.x - 5, 15, tar.z},
		tar,
		rl.Vector3{0, 1, 0},
		70.0,
		.PERSPECTIVE,
	}
}

handleCamera :: proc(g: ^Game) {
	// Zoom
	wheel := rl.GetMouseWheelMove()
	if wheel != 0 {
		rl.UpdateCameraPro(&g.camera, rl.Vector3(0), rl.Vector3(0), -wheel)
	}

	// Orbit
	if rl.IsMouseButtonDown(.MIDDLE) {
		delta := rl.GetMouseDelta()
		rotationSensitivity := f32(0.2)
		rl.UpdateCameraPro(
			&g.camera,
			rl.Vector3(0),
			rl.Vector3{delta.x * rotationSensitivity, delta.y * rotationSensitivity, 0},
			0.0,
		)
	}

	// Panning with Right Click? (Optional, good for UX)
	if rl.IsMouseButtonDown(.RIGHT) {
		delta := rl.GetMouseDelta()
		panSpeed := f32(0.05)
		rl.UpdateCameraPro(
			&g.camera,
			rl.Vector3{delta.y * panSpeed, -delta.x * panSpeed, 0},
			rl.Vector3(0),
			0.0,
		)
	}

	// Camera Shake
	if g.CameraShakeIntensity > 0 {
		g.CameraShakeIntensity -= 0.05 // Decay
		if g.CameraShakeIntensity < 0 {
			g.CameraShakeIntensity = 0
		}
	}
}

GetRenderCamera :: proc(g: ^Game) -> rl.Camera3D {
	cam := g.camera
	if g.CameraShakeIntensity > 0 {
		offset := rl.Vector3 {
			(f32(rl.GetRandomValue(0, 100)) / 50.0 - 1.0) * g.CameraShakeIntensity,
			(f32(rl.GetRandomValue(0, 100)) / 50.0 - 1.0) * g.CameraShakeIntensity,
			(f32(rl.GetRandomValue(0, 100)) / 50.0 - 1.0) * g.CameraShakeIntensity,
		}
		cam.position = cam.position + offset
		cam.target = cam.target + offset
	}
	return cam
}
