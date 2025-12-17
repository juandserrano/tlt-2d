package game

import (
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

type Tile struct {
	position   rl.Vector3
	model      *rl.Model
	gridX      int
	gridZ      int
	tileType   TileType
	isWalkable bool
	isOccupied bool
	isSpawn    bool
}

func (t *Tile) Draw() {
	rl.DrawModel(*t.model, t.position, 1, rl.White)
}
