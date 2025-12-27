package game

import "core:math"
import "core:math/rand"
import rl "vendor:raylib"

GridCoord :: struct {
	x: int,
	z: int,
}

// GridToWorldHex converts grid coordinates (col, row) to World Pixels
GridToWorldHex :: proc(col, row: int, size: f32) -> rl.Vector2 {
	// 1. Calculate dimensions based on Size (Flat-Top)
	// Width is 2 * size
	// hexWidth := 2.0 * size -- Unused
	// Height is sqrt(3) * size
	hexHeight := math.sqrt_f32(3) * size

	// Horizontal distance between centers (3/4 of width = 3/4 * 2 * size = 1.5 * size)
	horizDist := 1.5 * size
	// Vertical distance between centers (full height)
	vertDist := hexHeight

	// 2. Calculate X Position
	// Columns are offset by 3/4 width
	x := f32(col) * horizDist

	// 3. Calculate Y Position (mapped to Z in world)
	y := f32(row) * vertDist

	// OFFSET LOGIC (Odd-Q):
	// If the column is Odd, we shift this tile down (Postive Z/Y) by half a height
	// Go's % operator returns negative result for negative operands (-1 % 2 == -1), so != 0 is correct.
	if col % 2 != 0 {
		y += vertDist / 2.0
	}

	return rl.Vector2{x, y}
}

GenerateFlatTopGrid :: proc(countX, countZ: int, radius: f32) -> []Tile {
	arr: []Tile
	tiles: [dynamic]Tile

	// Flat-Top Hex Math
	// Width is distance from point to point (Horizontal)
	hexWidth := 2.0 * radius
	// Height is distance from flat side to flat side (Vertical)
	hexHeight := math.sqrt_f32(3) * radius

	// Distance between center points
	// We overlap on the X axis now
	horizDist := 0.75 * hexWidth
	vertDist := hexHeight

	startX := -countX / 2
	startZ := -countZ / 2

	for x := startX; x < startX + countX; x += 1 {
		for z := startZ; z < startZ + countZ; z += 1 {

			// Calculate World X
			xPos := f32(x) * horizDist

			// Calculate World Z
			zPos := f32(z) * vertDist

			// FLAT-TOP Specific: Offset every odd COLUMN (X)
			// We shift the Z position
			if x % 2 != 0 {
				zPos += vertDist / 2.0
			}

			pos := rl.Vector3{xPos, 0.0, zPos}

			tile := Tile {
				position = pos,
				gridX    = x,
				gridZ    = z,
				tileType = .TileTypeClear,
			}

			// // Modify spawnable tiles
			if tile.gridX == 12 && tile.gridZ == 0 ||
			   tile.gridX == -12 && tile.gridZ == 0 ||
			   tile.gridX == 6 && tile.gridZ == 9 ||
			   tile.gridX == -6 && tile.gridZ == 9 ||
			   tile.gridX == 6 && tile.gridZ == -9 ||
			   tile.gridX == -6 && tile.gridZ == -9 {
				tile.isSpawn = true
			}

			append(&tiles, tile)

			arr = tiles[:]

		}
	}

	return arr
}

GetTileCenter :: proc(g: ^Game, gPos: GridCoord) -> rl.Vector3 {
	for t in g.levels[g.currentLevel].tiles {
		if t.gridX == gPos.x && t.gridZ == gPos.z {
			return t.position
		}
	}

	return rl.Vector3(0)
}

GetTileWithGridPos :: proc(g: ^Game, gPos: GridCoord) -> ^Tile {
	for &t in g.levels[g.currentLevel].tiles {
		if t.gridX == gPos.x && t.gridZ == gPos.z {
			return &t
		}
	}

	return nil
}

GetAllSpawnableTileGridCoords :: proc(g: ^Game) -> []GridCoord {

	spawableTiles := []GridCoord{{12, 0}, {-12, 0}, {6, 9}, {-6, 9}, {-6, -9}, {6, -9}}

	return spawableTiles
}

GetRandomSpawnableTileGridCoord :: proc(g: ^Game) -> GridCoord {

	spawableTiles := []GridCoord{{12, 0}, {-12, 0}, {6, 9}, {-6, 9}, {-6, -9}, {6, -9}}

	return spawableTiles[rand.int_range(0, 7)]
}
