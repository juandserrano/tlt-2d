package game

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	tiles    []Tile
	xTiles   int
	zTiles   int
	id       int
	centerXZ rl.Vector2
}

func (g *Game) LoadLevel(level int) {
	var l Level
	l.xTiles = 40
	l.zTiles = 40

	// tiles := make([]Tile, l.xTiles*l.zTiles)

	// i := 0
	// for x := range l.xTiles {
	// 	for z := range l.zTiles {
	// 		tilePos := GridToWorldHex(x, z, HEX_TILE_WIDTH/2.0)
	// 		tiles[i].position.X = tilePos.X
	// 		tiles[i].position.Z = tilePos.Y

	// 		tiles[i].position.Y = -0.05 // lower tile so that 0 is top face
	// 		tiles[i].gridX = x
	// 		tiles[i].gridZ = z
	// 		tiles[i].model = g.tiles[TileTypeClear].model
	// 		// Apply model based on tile type
	// 		switch tiles[i].tileType {
	// 		case TileTypeClear:
	// 			tiles[i].model = g.tiles[TileTypeClear].model
	// 		default:
	// 			fmt.Println("Im drawing only clear")
	// 		}
	// 		i++
	// 	}
	// }

	// // Water tiles
	// tiles[5].tileType = TileTypeWater
	// tiles[5].model = g.tiles[TileTypeWater].model

	// tiles[6].tileType = TileTypeWater
	// tiles[6].model = g.tiles[TileTypeWater].model

	center := GridToWorldHex(l.xTiles/2, l.zTiles/2, HEX_TILE_WIDTH/2.0)

	newGrid := GenerateFlatTopGrid(l.xTiles, l.zTiles, HEX_TILE_WIDTH/2.0)
	for i := range newGrid {
		newGrid[i].model = g.tiles[TileTypeClear].model
		// Apply model based on tile type
		switch newGrid[i].tileType {
		case TileTypeClear:
			newGrid[i].model = g.tiles[TileTypeClear].model
		default:
			fmt.Println("Im drawing only clear")
		}
	}

	g.levels[level] = Level{
		id:       level,
		tiles:    newGrid,
		xTiles:   l.xTiles,
		zTiles:   l.zTiles,
		centerXZ: center,
	}
	g.currentLevel = level
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

func GenerateFlatTopGrid(countX, countZ int, radius float32) []Tile {
	var tiles []Tile

	// Flat-Top Hex Math
	// Width is distance from point to point (Horizontal)
	hexWidth := 2.0 * radius
	// Height is distance from flat side to flat side (Vertical)
	hexHeight := float32(math.Sqrt(3)) * radius

	// Distance between center points
	// We overlap on the X axis now
	horizDist := 0.75 * hexWidth
	vertDist := hexHeight

	startX := -countX / 2
	startZ := -countZ / 2

	for x := startX; x < startX+countX; x++ {
		for z := startZ; z < startZ+countZ; z++ {

			// Calculate World X
			xPos := float32(x) * horizDist

			// Calculate World Z
			zPos := float32(z) * vertDist

			// FLAT-TOP Specific: Offset every odd COLUMN (X)
			// We shift the Z position
			if x%2 != 0 {
				zPos += vertDist / 2.0
			}

			pos := rl.NewVector3(xPos, 0.0, zPos)

			tile := Tile{
				position: pos,
				gridX:    x,
				gridZ:    z,
				tileType: TileTypeClear,
			}

			tiles = append(tiles, tile)
		}
	}

	return tiles
}
