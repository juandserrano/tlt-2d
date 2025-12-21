package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Castle struct {
	model    rl.Model
	position rl.Vector3
	gridX    int
	gridZ    int
	health   int
}

func (p *Castle) draw() {
	rl.DrawModel(p.model, p.position, 1.0, rl.White)
}

func (g *Game) initPlayerCastle() {
	g.playerCastle.gridX = 0
	g.playerCastle.gridZ = 0
	startXZPos := GridToWorldHex(g.playerCastle.gridX, g.playerCastle.gridZ, HEX_TILE_WIDTH/2.0)
	g.playerCastle.position.X = startXZPos.X
	g.playerCastle.position.Z = startXZPos.Y
	g.playerCastle.position.Y = 0
	g.playerCastle.health = g.Config.Player.Health
}

func (g *Game) TurnPlayer(dt float32) {
	if len(g.playerHand.cards) < g.playerHand.maxCards {
		g.UI.buttons["draw"].enabled = true
	}
	g.UI.buttons["end_turn"].enabled = true
	g.handlePlayingInput(dt)
	g.highlightValidEnemies()

}

func (g *Game) highlightValidEnemies() {
	if g.playerHand.selectedCard == nil {
		return
	}

	switch g.playerHand.selectedCard.cardType {
	case CardTypeAttackPawn:
		for i := range EnemiesInPlay {
			if EnemiesInPlay[i].enemyType == EnemyTypePawn {
				EnemiesInPlay[i].isHighlighted = true
			}
		}
	case CardTypeAttackKnight:
		for i := range EnemiesInPlay {
			if EnemiesInPlay[i].enemyType == EnemyTypeKnight {
				EnemiesInPlay[i].isHighlighted = true
			}
		}
	case CardTypeAttackBishop:
		for i := range EnemiesInPlay {
			if EnemiesInPlay[i].enemyType == EnemyTypeBishop {
				EnemiesInPlay[i].isHighlighted = true
			}
		}

	}
}
