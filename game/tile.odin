package game

import rl "vendor:raylib"
HEX_TILE_WIDTH :: 1.16
HEX_TILE_SIZE :: HEX_TILE_WIDTH / 2.0

TileType :: enum {
	TileTypeClear,
	TileTypeGrass,
	TileTypeDirt,
	TileTypeMountain,
	TileTypeRocks,
	TileTypeWater,
}

Tile :: struct {
	position:   rl.Vector3,
	model:      ^rl.Model,
	gridX:      int,
	gridZ:      int,
	tileType:   TileType,
	isWalkable: bool,
	isOccupied: bool,
	isSpawn:    bool,
}

drawTile :: proc(t: ^Tile) {
	rl.DrawModel(t.model^, t.position, 1, rl.WHITE)
}
