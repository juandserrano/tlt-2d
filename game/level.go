package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	tiles  []Tile
	xTiles int
	zTiles int
	id     int
}

func (g *Game) InitLevel1() {
	l1 := g.levels[1]
	l1.xTiles = 10
	l1.zTiles = 3
	bbox := rl.GetModelBoundingBox(g.basicTileModel)
	size := rl.Vector3Subtract(bbox.Max, bbox.Min)

	tiles := make([]Tile, l1.xTiles*l1.zTiles)

	i := 0
	for x := range l1.xTiles {
		for z := range l1.zTiles {
			tiles[i].width = size.X
			tiles[i].length = size.Z
			tiles[i].height = size.Y
			tiles[i].position.X = float32(x) * (tiles[i].width)
			tiles[i].position.Z = float32(z) * (tiles[i].length)
			tiles[i].position.Y = -0.05 // lower tile so that 0 is top face
			tiles[i].x = x
			tiles[i].z = z
			tiles[i].model = g.basicTileModel
			i++
		}
	}

	g.levels[1] = Level{
		id:    1,
		tiles: tiles,
	}
}

func (l *Level) Draw() {
	for _, t := range l.tiles {
		t.Draw()
	}
}
