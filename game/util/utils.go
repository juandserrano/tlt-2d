package util

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var hitpos rl.Vector3

// GetHexFromMouse returns the grid coordinates (q, r) of the hex under the mouse.
// Returns (0, 0, false) if the mouse clicked the sky (didn't hit the ground plane).
func GetHexFromMouse(camera rl.Camera3D, hexSize float32) (int, int, bool) {
	// 1. Get the Ray from the Camera to the Mouse
	ray := rl.GetScreenToWorldRay(rl.GetMousePosition(), camera)

	// 2. Intersect Ray with the Ground Plane
	// We assume our grid is flat on the XZ plane (Y = 0)
	// Plane Normal points UP (0, 1, 0)
	// Plane Center is at (0, 0, 0)
	groundHitInfo := rl.GetRayCollisionQuad(
		ray,
		rl.NewVector3(-1000, 0, -1000), // Large quad corners
		rl.NewVector3(-1000, 0, 1000),
		rl.NewVector3(1000, 0, 1000),
		rl.NewVector3(1000, 0, -1000),
	)

	if !groundHitInfo.Hit {
		return 0, 0, false // Clicked off the map/into the sky
	}

	// 3. Convert the Hit Point (3D World) to Hex Grid (2D Logical)
	// Note: We map World X -> Hex X, and World Z -> Hex Y
	hitpos = groundHitInfo.Point
	fmt.Println("world x:", hitpos.X, "- world z:", hitpos.Z)

	// Convert world pixels to hex axial coords
	q, r := WorldToHex(hitpos.X, hitpos.Z, hexSize)

	return q, r, true
}

// WorldToHex handles the math of converting 3D coordinates to Hex Grid IDs
// WorldToHex handles the math of converting 3D coordinates to Hex Grid IDs
func WorldToHex(x, z float32, size float32) (int, int) {
	// 1. Axial Coordinate Math (Flat-Top Hex)
	// q = (2/3 * x) / size
	// r = (-1/3 * x + sqrt(3)/3 * z) / size

	qFloat := (2.0 / 3.0 * float64(x)) / float64(size)
	rFloat := (-1.0/3.0*float64(x) + math.Sqrt(3)/3.0*float64(z)) / float64(size)

	// 2. Rounding
	return AxialRoundToOffset(qFloat, rFloat)
}

// Helper: Round fractional hex coords to the nearest valid integer hex
func AxialRoundToOffset(frQ, frR float64) (int, int) {
	x, z := frQ, frR
	y := -x - z

	rx := math.Round(x)
	ry := math.Round(y)
	rz := math.Round(z)

	diffX := math.Abs(rx - x)
	diffY := math.Abs(ry - y)
	diffZ := math.Abs(rz - z)

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
	row := int(rz + (rx-float64(int(rx)&1))/2.0)

	return col, row
}

func GetDirectionTo(from, to rl.Vector3) rl.Vector3 {
	return rl.Vector3Normalize(rl.Vector3Subtract(from, to))
}

func CalculateRotation(from, to rl.Vector3) int {

	dir := GetDirectionTo(from, to)
	angleRad := math.Atan2(float64(dir.Z), float64(dir.X))

	angleDeg := float32(angleRad * (180.0 / math.Pi))
	angleDeg -= 90.0
	return int(angleDeg)
}
