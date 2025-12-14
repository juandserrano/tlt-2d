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
	// l := g.levels[level]
	var l Level
	l.xTiles = 20
	l.zTiles = 3

	tiles := make([]Tile, l.xTiles*l.zTiles)

	i := 0
	for x := range l.xTiles {
		for z := range l.zTiles {
			tilePos := GridToWorldHex(x, z, HEX_TILE_WIDTH/2.0)
			tiles[i].position.X = tilePos.X
			tiles[i].position.Z = tilePos.Y

			tiles[i].position.Y = -0.05 // lower tile so that 0 is top face
			tiles[i].x = x
			tiles[i].z = z
			tiles[i].model = g.tiles[TileTypeClear].model
			// Apply model based on tile type
			switch tiles[i].tileType {
			case TileTypeClear:
				tiles[i].model = g.tiles[TileTypeClear].model
			default:
				fmt.Println("Im drawing only clear")
			}
			i++
		}
	}

	center := GridToWorldHex(l.xTiles/2, l.zTiles/2, HEX_TILE_WIDTH/2.0)

	g.levels[level] = Level{
		id:       level,
		tiles:    tiles,
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

	return rl.Vector2{X: y, Y: x}
}
