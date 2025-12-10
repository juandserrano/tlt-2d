package game

import rl "github.com/gen2brain/raylib-go/raylib"

const GRASS_TILE_SCALE = 0.1

type Tile struct {
	position rl.Vector2
	texture  rl.Texture2D
}

func (t *Tile) Draw() {
	rl.DrawTextureEx(t.texture, rl.NewVector2(t.position.X+400, t.position.Y+100), 0, GRASS_TILE_SCALE, rl.White)
}
