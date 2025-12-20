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
	attackPawnCardTexture := rl.LoadTexture("assets/textures/cards/attack_pawn.png")
	attackKnightCardTexture := rl.LoadTexture("assets/textures/cards/attack_knight.png")
	attackBishopCardTexture := rl.LoadTexture("assets/textures/cards/attack_bishop.png")
	backCardTexture := rl.LoadTexture("assets/textures/cards/card_back.png")
	g.enemyModels[EnemyTypePawn] = &pawnModel
	g.enemyModels[EnemyTypeKnight] = &knightModel
	g.enemyModels[EnemyTypeBishop] = &bishopModel
	g.cardTextures[CardTypeAttackPawn] = &attackPawnCardTexture
	g.cardTextures[CardTypeAttackKnight] = &attackKnightCardTexture
	g.cardTextures[CardTypeAttackBishop] = &attackBishopCardTexture
	g.cardTextures[CardTypeBack] = &backCardTexture

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
