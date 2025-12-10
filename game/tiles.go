package game

import rl "github.com/gen2brain/raylib-go/raylib"

const GRASS_TILE_SCALE = 0.1

type Tile struct {
	position rl.Vector2
	texture  rl.Texture2D
	x        int
	y        int
}

func (t *Tile) Draw() {
	rl.DrawTextureEx(t.texture, rl.NewVector2(t.position.X-(float32(t.texture.Width)*GRASS_TILE_SCALE)/2, t.position.Y-float32(t.texture.Height)*GRASS_TILE_SCALE/2), 0, GRASS_TILE_SCALE, rl.White)
}
