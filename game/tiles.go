package game

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const HEX_TILE_WIDTH = 1.16

type TileType int

const (
	TileTypeClear TileType = iota
	TileTypeGrass
	TileTypeDirt
	TileTypeMountain
	TileTypeRocks
	TileTypeWater
)

var hitpos rl.Vector3

type Tile struct {
	position rl.Vector3
	model    *rl.Model
	gridX    int
	gridZ    int
	tileType TileType
}

func (t *Tile) Draw() {
	rl.DrawModel(*t.model, t.position, 1, rl.White)
}

func (t *Tile) debugDrawGridCoord() {
	// --- DRAW TEXT 3D ---
	// 1. Push the current matrix so we don't mess up other 3D objects
	rl.PushMatrix()

	// 2. Move to the position in 3D space (X, Y, Z)
	rl.Translatef(t.position.X, 0.1, t.position.Z)

	// 3. Rotate the text.
	// By default, text lies flat on the floor looking up.
	// Rotate 90 degrees on X to make it stand up.
	// Rotate 180 degrees on Y because text usually renders "backwards" in 3D look-at logic.
	rl.Rotatef(90, 1, 0, 0)
	// rl.Rotatef(90, 0, 1, 0)
	rl.Rotatef(90, 0, 0, 1)

	// 4. Scale it DOWN.
	// Standard font size 20 is "20 meters" high in 3D.
	// We scale by 0.1 to make it manageable.
	rl.Scalef(0.01, 0.01, 0.01)

	// 5. Draw the text (Relative to 0,0 because we already translated the matrix)
	// We center the text by calculating width/2
	text := fmt.Sprintf("(%d, %d)", t.gridX, t.gridZ)
	fontSize := int32(20)
	textWidth := rl.MeasureText(text, fontSize)

	rl.DrawText(text, -textWidth/2, 0, fontSize, rl.Red)

	// 6. Restore the matrix
	rl.PopMatrix()
	// --------------------

}

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
func WorldToHex(x, z float32, size float32) (int, int) {
	// 1. Axial Coordinate Math (Inverted Pointy-Top Hex)
	// x_axial = (sqrt(3)/3 * x  -  1/3 * z) / size
	// z_axial = (2/3 * z) / size

	qFloat := (float64(x)*math.Sqrt(3)/3.0 - float64(z)/3.0) / float64(size)
	rFloat := (float64(z) * 2.0 / 3.0) / float64(size)

	// 2. Rounding (The crucial step)
	// We can't just cast to int because of the jagged edges
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
		ry = -rx - rz
	} else {
		rz = -rx - ry
	}

	// Convert Cubic (x, z, y) back to Offset (Odd-R)
	// Note: Our "WorldToHex" output 'q' maps to x, 'r' maps to z (row)
	col := int(rx + (rz-float64(int(rz)&1))/2.0)
	row := int(rz)

	return col, row
}
