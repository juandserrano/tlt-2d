package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) LoadResources() {
	g.LoadTileModels()
	g.LoadTextures()
}

func (g *Game) LoadTileModels() {
	g.basicTileModel = rl.LoadModel("assets/models/tiles/basic_ground_tile.glb")

}

func (g *Game) LoadTextures() {

}
