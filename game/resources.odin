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
	g.enemyModels[.EnemyTypePawn] = rl.LoadModel("assets/models/enemies/pawn.glb")
	g.enemyModels[.EnemyTypeKnight] = rl.LoadModel("assets/models/enemies/knight.glb")
	g.enemyModels[.EnemyTypeBishop] = rl.LoadModel("assets/models/enemies/bishop.glb")
	g.enemyModels[.EnemyTypeQueen] = rl.LoadModel("assets/models/enemies/queen.glb")
	g.enemyModels[.EnemyTypeKing] = rl.LoadModel("assets/models/enemies/king.glb")
	g.cardTextures[.CardTypeAttackPawn] = rl.LoadTexture("assets/textures/cards/attack_pawn.png")
	g.cardTextures[.CardTypeAttackKnight] = rl.LoadTexture(
		"assets/textures/cards/attack_knight.png",
	)
	g.cardTextures[.CardTypeAttackBishop] = rl.LoadTexture(
		"assets/textures/cards/attack_bishop.png",
	)
	g.cardTextures[.CardTypeAttackQueen] = rl.LoadTexture("assets/textures/cards/attack_queen.png")
	g.cardTextures[.CardTypeAttackKing] = rl.LoadTexture("assets/textures/cards/attack_king.png")
	g.cardTextures[.CardTypeBack] = rl.LoadTexture("assets/textures/cards/card_back.png")

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
