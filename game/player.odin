package game

import rl "vendor:raylib"

Castle :: struct {
	model:    rl.Model,
	position: rl.Vector3,
	gridX:    int,
	gridZ:    int,
	health:   int,
}

drawCastle :: proc(p: ^Castle) {
	rl.DrawModel(p.model, p.position, 1.0, rl.WHITE)
}

initPlayerCastle :: proc(g: ^Game) {
	g.playerCastle.gridX = 0
	g.playerCastle.gridZ = 0
	startXZPos := GridToWorldHex(g.playerCastle.gridX, g.playerCastle.gridZ, HEX_TILE_WIDTH / 2.0)
	g.playerCastle.position.x = startXZPos.x
	g.playerCastle.position.z = startXZPos.y
	g.playerCastle.position.y = 0
	g.playerCastle.health = g.Config.Player.Health
}

TurnPlayer :: proc(g: ^Game, dt: f32) {
	// if g.deck.canDraw {
	// 	g.UI.buttons["draw"].enabled = true
	// } else {
	// 	g.UI.buttons["draw"].enabled = false
	// }
	// g.UI.buttons["end_turn"].enabled = true
	// g.handlePlayingInput(dt)
	// g.highlightValidEnemies()

}

highlightValidEnemies :: proc(g: ^Game) {
	// if g.playerHand.selectedCard == nil {
	// 	return
	// }

	// switch g.playerHand.selectedCard.cardType {
	// case CardTypeAttackPawn:
	// 	for i in EnemiesInPlay {
	// 		if EnemiesInPlay[i].enemyType == EnemyTypePawn {
	// 			EnemiesInPlay[i].isHighlighted = true
	// 		}
	// 	}
	// case CardTypeAttackKnight:
	// 	for i in EnemiesInPlay {
	// 		if EnemiesInPlay[i].enemyType == EnemyTypeKnight {
	// 			EnemiesInPlay[i].isHighlighted = true
	// 		}
	// 	}
	// case CardTypeAttackBishop:
	// 	for i in EnemiesInPlay {
	// 		if EnemiesInPlay[i].enemyType == EnemyTypeBishop {
	// 			EnemiesInPlay[i].isHighlighted = true
	// 		}
	// 	}
	// case CardTypeAttackQueen:
	// 	for i in EnemiesInPlay {
	// 		if EnemiesInPlay[i].enemyType == EnemyTypeQueen {
	// 			EnemiesInPlay[i].isHighlighted = true
	// 		}
	// 	}
	// case CardTypeAttackKing:
	// 	for i in EnemiesInPlay {
	// 		if EnemiesInPlay[i].enemyType == EnemyTypeKing {
	// 			EnemiesInPlay[i].isHighlighted = true
	// 		}
	// 	}
	// }
}
