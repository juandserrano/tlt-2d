package game

import "core:math"
import rl "vendor:raylib"

Level :: struct {
	tiles:    []Tile,
	xTiles:   int,
	zTiles:   int,
	id:       int,
	centerXZ: rl.Vector2,
}

LoadLevelTiles :: proc(g: ^Game, level: int) {
	l: Level
	l.xTiles = 30
	l.zTiles = 30

	// // Water tiles
	// tiles[5].tileType = TileTypeWater
	// tiles[5].model = g.tiles[TileTypeWater].model

	// tiles[6].tileType = TileTypeWater
	// tiles[6].model = g.tiles[TileTypeWater].model

	center := GridToWorldHex(l.xTiles / 2, l.zTiles / 2, HEX_TILE_WIDTH / 2.0)

	newGrid := GenerateFlatTopGrid(l.xTiles, l.zTiles, HEX_TILE_WIDTH / 2.0)
	for &t in newGrid {
		t.model = g.tiles[.TileTypeClear].model
		// Apply model based on tile type
		#partial switch t.tileType {
		case .TileTypeClear:
			t.model = g.tiles[.TileTypeClear].model
		case:
		}
		// Assign ambient shader to tile models
		// materials := t.model.materials
		// for m in materials {
		// 	m.Shader = g.shaders[.AmbientShader]
		// }
	}

	g.levels[level] = Level {
		id       = level,
		tiles    = newGrid,
		xTiles   = l.xTiles,
		zTiles   = l.zTiles,
		centerXZ = center,
	}
	g.currentLevel = level
}

drawLevel :: proc(l: ^Level) {
	for &t in l.tiles {
		drawTile(&t)
	}
}
