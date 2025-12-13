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

//func (l *Level) ySortTiles() {
//	sort.Slice(l.tiles[:], func(i, j int) bool {
//		return l.tiles[i].position.Y < l.tiles[j].position.Y
//	})
//}

// func ySortTiles(tiles []Tile) {
// 	sort.Slice(tiles[:], func(i, j int) bool {
// 		return tiles[i].position.Y < tiles[j].position.Y
// 	})
// }

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
			// tiles[i].texture = g.tex
			tiles[i].model = g.basicTileModel
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
