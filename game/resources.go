package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) LoadResources() {
	g.LoadModels()
	g.LoadTextures()
}

func (g *Game) LoadModels() {
	g.plainTileModel = rl.LoadModel("assets/models/tiles/basic_ground_tile.glb")
	g.waterTileModel = rl.LoadModel("assets/models/tiles/basic_ground_tile.glb")
	g.pawnModel = rl.LoadModel("assets/models/enemies/pawn.glb")
	g.knightModel = rl.LoadModel("assets/models/enemies/knight.glb")
	g.playerCastle.model = rl.LoadModel("assets/models/castle/tower.glb")

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
