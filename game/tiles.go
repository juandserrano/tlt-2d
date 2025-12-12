package game

import rl "github.com/gen2brain/raylib-go/raylib"

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
	// rl.DrawTextureEx(t.texture, rl.NewVector2(t.position.X-(float32(t.texture.Width)*GRASS_TILE_SCALE)/2, t.position.Y-float32(t.texture.Height)*GRASS_TILE_SCALE/2), 0, GRASS_TILE_SCALE, rl.White)
}

// func (l *Level) GetTileCenterPosition(x, y int) rl.Vector2 {
// 	for _, t := range l.tiles {
// 		if t.x == x && t.y == y {
// 			return t.position
// 		}
// 	}
// 	return rl.Vector2Zero()
// }

func (g *Game) LoadBasicTile() {
	g.basicTileModel = rl.LoadModel("assets/plain_tile.obj")
}
