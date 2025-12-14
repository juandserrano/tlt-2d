package game

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	tiles  []Tile
	xTiles int
	zTiles int
	id     int
}

func (g *Game) InitLevel1() {
	l1 := g.levels[1]
	l1.xTiles = 10
	l1.zTiles = 3
	bbox := rl.GetModelBoundingBox(g.basicTileModel)
	size := rl.Vector3Subtract(bbox.Max, bbox.Min)

	tiles := make([]Tile, l1.xTiles*l1.zTiles)

	i := 0
	for x := range l1.xTiles {
		for z := range l1.zTiles {
			tiles[i].width = size.X
			tiles[i].length = size.Z
			tiles[i].height = size.Y
			tilePos := GridToWorldHex(x, z, 0.5)
			tiles[i].position.X = tilePos.X
			tiles[i].position.Z = tilePos.Y

			// tiles[i].position.X = float32(x) * (1)                // * (tiles[i].width)
			// tiles[i].position.Z = float32(z) * (3.0 / 4.0 * 1.16) // * (tiles[i].length)
			tiles[i].position.Y = -0.05 // lower tile so that 0 is top face
			tiles[i].x = x
			tiles[i].z = z
			tiles[i].model = g.basicTileModel
			i++
		}
	}

	g.levels[1] = Level{
		id:    1,
		tiles: tiles,
	}
}

func (l *Level) Draw() {
	for _, t := range l.tiles {
		t.Draw()
	}
}

// GridToWorldHex converts grid coordinates (col, row) to World Pixels
func GridToWorldHex(col, row int, size float32) rl.Vector2 {
	// 1. Calculate dimensions based on Size
	// Width of a pointy hex is sqrt(3) * size
	hexWidth := float32(math.Sqrt(3)) * size

	// Height is 2 * size, but rows overlap by 1/4, so vertical step is 1.5 * size
	vertDist := size * 1.5

	// 2. Calculate X Position
	// Standard X spacing is hexWidth
	x := float32(col) * hexWidth

	// OFFSET LOGIC:
	// If the row is Odd, we shift this tile right by half a width
	if row%2 == 1 {
		x += hexWidth / 2.0
	}

	// 3. Calculate Y Position
	y := float32(row) * vertDist

	return rl.Vector2{X: x, Y: y}
}
