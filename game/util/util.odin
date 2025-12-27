package util

import "core:fmt"
import "core:math"
import rl "vendor:raylib"

hitpos: rl.Vector3

// GetHexFromMouse returns the grid coordinates (q, r) of the hex under the mouse.
// Returns (0, 0, false) if the mouse clicked the sky (didn't hit the ground plane).
GetHexFromMouse :: proc(camera: rl.Camera3D, hexSize: f32) -> (int, int, bool) {
	// 1. Get the Ray from the Camera to the Mouse
	ray := rl.GetScreenToWorldRay(rl.GetMousePosition(), camera)

	// 2. Intersect Ray with the Ground Plane
	// We assume our grid is flat on the XZ plane (Y = 0)
	// Plane Normal points UP (0, 1, 0)
	// Plane Center is at (0, 0, 0)
	groundHitInfo := rl.GetRayCollisionQuad(
	ray,
	rl.Vector3{-1000, 0, -1000}, // Large quad corners
	rl.Vector3{-1000, 0, 1000},
	rl.Vector3{1000, 0, 1000},
	rl.Vector3{1000, 0, -1000},
	)

	if !groundHitInfo.hit {
		return 0, 0, false // Clicked off the map/into the sky
	}

	// 3. Convert the Hit Point (3D World) to Hex Grid (2D Logical)
	// Note: We map World X -> Hex X, and World Z -> Hex Y
	hitpos = groundHitInfo.point
	fmt.println("world x:", hitpos.x, "- world z:", hitpos.z)

	// Convert world pixels to hex axial coords
	q, r := WorldToHex(hitpos.x, hitpos.z, hexSize)

	return q, r, true
}

// WorldToHex handles the math of converting 3D coordinates to Hex Grid IDs
// WorldToHex handles the math of converting 3D coordinates to Hex Grid IDs
WorldToHex :: proc(x, z, size: f32) -> (int, int) {
	// 1. Axial Coordinate Math (Flat-Top Hex)
	// q = (2/3 * x) / size
	// r = (-1/3 * x + sqrt(3)/3 * z) / size

	qFloat := (2.0 / 3.0 * f64(x)) / f64(size)
	rFloat := (-1.0 / 3.0 * f64(x) + math.sqrt_f64(3) / 3.0 * f64(z)) / f64(size)

	// 2. Rounding
	return AxialRoundToOffset(qFloat, rFloat)
}

// Helper: Round fractional hex coords to the nearest valid integer hex
AxialRoundToOffset :: proc(frQ, frR: f64) -> (int, int) {
	x, z := frQ, frR
	y := -x - z

	rx := math.round(x)
	ry := math.round(y)
	rz := math.round(z)

	diffX := math.abs(rx - x)
	diffY := math.abs(ry - y)
	diffZ := math.abs(rz - z)

	if diffX > diffY && diffX > diffZ {
		rx = -ry - rz
	} else if diffY > diffZ {
		// ry = -rx - rz
	} else {
		rz = -rx - ry
	}

	// Convert Cubic (q, r, s) to Offset (Odd-Q)
	// q = rx, r = rz
	// col = q
	// row = r + (q - (q&1)) / 2
	col := int(rx)
	row := int(rz + (rx - f64(int(rx) & 1)) / 2.0)

	return col, row
}
GetDirectionTo :: proc(from, to: rl.Vector3) -> rl.Vector3 {
	return rl.Vector3Normalize(from - to)
}

CalculateRotation :: proc(from, to: rl.Vector3) -> int {

	dir := GetDirectionTo(from, to)
	angleRad := math.atan2(f64(dir.z), f64(dir.x))

	angleDeg := f32(angleRad * (180.0 / math.PI))
	angleDeg -= 90.0
	return int(angleDeg)
}
