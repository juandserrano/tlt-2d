package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) LoadResources() {
	g.LoadModels()
	g.LoadTextures()
}

func (g *Game) LoadModels() {
	g.playerCastle.model = rl.LoadModel("assets/models/castle/tower.glb")
	g.plainTileModel = rl.LoadModel("assets/models/tiles/basic_ground_tile.glb")
	g.waterTileModel = rl.LoadModel("assets/models/tiles/basic_ground_tile.glb")
	pawnModel := rl.LoadModel("assets/models/enemies/pawn.glb")
	knightModel := rl.LoadModel("assets/models/enemies/knight.glb")
	bishopModel := rl.LoadModel("assets/models/enemies/bishop.glb")
	// attackPawnCardModel := rl.LoadModel("assets/models/cards/attack_pawn.glb")
	attackPawnCardTexture := rl.LoadTexture("assets/textures/cards/attack_pawn.png")
	g.enemyModels[EnemyTypePawn] = &pawnModel
	g.enemyModels[EnemyTypeKnight] = &knightModel
	g.enemyModels[EnemyTypeBishop] = &bishopModel
	// g.cardModels[CardTypeAttackPawn] = &attackPawnCardModel
	g.cardTextures[CardTypeAttackPawn] = &attackPawnCardTexture

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
