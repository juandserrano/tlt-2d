package game

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GridCoord struct {
	X int
	Z int
}
type Level struct {
	tiles    []Tile
	xTiles   int
	zTiles   int
	id       int
	centerXZ rl.Vector2
}

func (g *Game) LoadLevelTiles(level int) {
	var l Level
	l.xTiles = 50
	l.zTiles = 50

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
		// Assign ambient shader to tile models
		materials := newGrid[i].model.GetMaterials()
		for i := range materials {
			materials[i].Shader = g.ambientShader
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
	// 1. Calculate dimensions based on Size (Flat-Top)
	// Width is 2 * size
	// hexWidth := 2.0 * size -- Unused
	// Height is sqrt(3) * size
	hexHeight := float32(math.Sqrt(3)) * size

	// Horizontal distance between centers (3/4 of width = 3/4 * 2 * size = 1.5 * size)
	horizDist := 1.5 * size
	// Vertical distance between centers (full height)
	vertDist := hexHeight

	// 2. Calculate X Position
	// Columns are offset by 3/4 width
	x := float32(col) * horizDist

	// 3. Calculate Y Position (mapped to Z in world)
	y := float32(row) * vertDist

	// OFFSET LOGIC (Odd-Q):
	// If the column is Odd, we shift this tile down (Postive Z/Y) by half a height
	// Go's % operator returns negative result for negative operands (-1 % 2 == -1), so != 0 is correct.
	if col%2 != 0 {
		y += vertDist / 2.0
	}

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

			// // Modify spawnable tiles
			if tile.gridX == 12 && tile.gridZ == 0 ||
				tile.gridX == -12 && tile.gridZ == 0 ||
				tile.gridX == 6 && tile.gridZ == 10 ||
				tile.gridX == -6 && tile.gridZ == 10 ||
				tile.gridX == 6 && tile.gridZ == -10 ||
				tile.gridX == -6 && tile.gridZ == -10 {
				tile.isSpawn = true
			}

			tiles = append(tiles, tile)

		}
	}

	return tiles
}

func (g *Game) GetTileCenter(gPos GridCoord) rl.Vector3 {
	for _, t := range g.levels[g.currentLevel].tiles {
		if t.gridX == gPos.X && t.gridZ == gPos.Z {
			return t.position
		}
	}

	return rl.Vector3Zero()
}

func (g *Game) GetTileWithGridPos(gPos GridCoord) *Tile {
	for _, t := range g.levels[g.currentLevel].tiles {
		if t.gridX == gPos.X && t.gridZ == gPos.Z {
			return &t
		}
	}

	return nil
}
