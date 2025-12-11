package game

import (
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	tiles  []Tile
	xTiles int
	yTiles int
	id     int
}

//func (l *Level) ySortTiles() {
//	sort.Slice(l.tiles[:], func(i, j int) bool {
//		return l.tiles[i].position.Y < l.tiles[j].position.Y
//	})
//}

func ySortTiles(tiles []Tile) {
	sort.Slice(tiles[:], func(i, j int) bool {
		return tiles[i].position.Y < tiles[j].position.Y
	})
}

func (g *Game) InitLevel1() {
	l1 := g.levels[1]
	l1.xTiles = 200
	l1.yTiles = 300

	tiles := make([]Tile, l1.xTiles*l1.yTiles)

	i := 0
	for x := range l1.xTiles {
		for y := range l1.yTiles {
			tiles[i].position.X = (float32(x) - float32(y)) * ((float32(g.tex.Width) / 2) * GRASS_TILE_SCALE)
			tiles[i].position.Y = (float32(x) + float32(y)) * ((float32(g.tex.Height)/2 - 50) * GRASS_TILE_SCALE)
			tiles[i].x = x
			tiles[i].y = y
			tiles[i].texture = g.tex
			i++
		}
	}

	//ySortTiles(tiles)

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

// GetTileCenter returns the visual center point of the tile.
// Useful for centering the camera or placing a unit standing in the middle.
func GetTileCenter(x, y int, tileTexture rl.Texture2D) rl.Vector2 {

	// 2. Calculate Center Grid Coordinate (e.g., 5.5, 5.5)
	centerX := float32(x) // + 0.5
	centerY := float32(y) // + 0.5

	// 3. Isometric Math for the center
	isoX := (centerX - centerY) * (float32(tileTexture.Width) * GRASS_TILE_SCALE / 2)
	isoY := (centerX + centerY) * (float32(tileTexture.Height) * GRASS_TILE_SCALE / 2)

	return rl.Vector2{X: isoX, Y: isoY}
}
