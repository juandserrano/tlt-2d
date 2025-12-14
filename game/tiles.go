package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const GRASS_TILE_SCALE = 0.1

type Tile struct {
	position rl.Vector3
	model    rl.Model
	x        int
	z        int
	width    float32
	length   float32
	height   float32
}

func (t *Tile) Draw() {
	rl.DrawModel(t.model, t.position, 1, rl.White)
}

func (g *Game) LoadBasicTile() {
	g.basicTileModel = rl.LoadModel("assets/models/tiles/basic_ground_tile.glb")

}
