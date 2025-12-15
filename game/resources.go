package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) LoadResources() {
	g.LoadTileModels()
	g.LoadTextures()
}

func (g *Game) LoadTileModels() {
	g.plainTileModel = rl.LoadModel("assets/models/tiles/basic_ground_tile.glb")
	g.waterTileModel = rl.LoadModel("assets/models/tiles/basic_ground_tile.glb")

	g.tiles[TileTypeClear] = Tile{
		model:    &g.plainTileModel,
		tileType: TileTypeClear,
	}
	g.tiles[TileTypeWater] = Tile{
		model:    &g.waterTileModel,
		tileType: TileTypeWater,
	}

}

func (g *Game) LoadTextures() {

}
