package game

import "core:fmt"
import rl "vendor:raylib"

LoadResources :: proc(g: ^Game) {
	// g.ParticleManager = NewParticleManager(2000)
	// g.ParticleManager.Init()

	LoadModels(g)
	LoadTextures(g)
	LoadSounds(g)
}

LoadSounds :: proc(g: ^Game) {
	g.sounds["chess_slide"] = rl.LoadSound("assets/sounds/chess_slide.wav")
	g.sounds["falling_impact"] = rl.LoadSound("assets/sounds/falling_impact.wav")
	g.sounds["card_select"] = rl.LoadSound("assets/sounds/microtick.wav")

	g.music["iron_at_the_gate"] = rl.LoadMusicStream("assets/sounds/iron_at_the_gate.mp3")
}

LoadModels :: proc(g: ^Game) {
	g.playerCastle.model = rl.LoadModel("assets/models/castle/tower.glb")
	g.plainTileModel = rl.LoadModel("assets/models/tiles/basic_ground_tile.glb")
	g.waterTileModel = rl.LoadModel("assets/models/tiles/basic_ground_tile.glb")
	pawnModel := rl.LoadModel("assets/models/enemies/pawn.glb")
	knightModel := rl.LoadModel("assets/models/enemies/knight.glb")
	bishopModel := rl.LoadModel("assets/models/enemies/bishop.glb")
	kingModel := rl.LoadModel("assets/models/enemies/king.glb")
	queenModel := rl.LoadModel("assets/models/enemies/queen.glb")
	attackPawnCardTexture := rl.LoadTexture("assets/textures/cards/attack_pawn.png")
	attackKnightCardTexture := rl.LoadTexture("assets/textures/cards/attack_knight.png")
	attackBishopCardTexture := rl.LoadTexture("assets/textures/cards/attack_bishop.png")
	attackQueenCardTexture := rl.LoadTexture("assets/textures/cards/attack_queen.png")
	attackKingCardTexture := rl.LoadTexture("assets/textures/cards/attack_king.png")
	backCardTexture := rl.LoadTexture("assets/textures/cards/card_back.png")
	g.enemyModels[.EnemyTypePawn] = &pawnModel
	g.enemyModels[.EnemyTypeKnight] = &knightModel
	g.enemyModels[.EnemyTypeBishop] = &bishopModel
	g.enemyModels[.EnemyTypeQueen] = &queenModel
	g.enemyModels[.EnemyTypeKing] = &kingModel
	g.cardTextures[.CardTypeAttackPawn] = &attackPawnCardTexture
	g.cardTextures[.CardTypeAttackKnight] = &attackKnightCardTexture
	g.cardTextures[.CardTypeAttackBishop] = &attackBishopCardTexture
	g.cardTextures[.CardTypeAttackQueen] = &attackQueenCardTexture
	g.cardTextures[.CardTypeAttackKing] = &attackKingCardTexture
	g.cardTextures[.CardTypeBack] = &backCardTexture

	g.tiles[.TileTypeClear] = Tile {
		model    = &g.plainTileModel,
		tileType = .TileTypeClear,
	}
	g.tiles[.TileTypeWater] = Tile {
		model    = &g.waterTileModel,
		tileType = .TileTypeWater,
	}

}

LoadTextures :: proc(g: ^Game) {

}
