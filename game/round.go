package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Round struct {
	TurnNumber int
}

func (g *Game) NewRound() Round {
	return Round{
		TurnNumber: 0,
	}
}

func (r *Round) SetUp(g *Game) {
	g.enemyBag = g.NewEnemyBag()
	g.playerHand = g.NewHand()
	g.discardPile = g.NewDiscardPile()
	g.deck = g.NewDeck()
	g.UI.buttons["draw"] = NewButton("draw", 300, 100, func() { g.drawToTopHand(&g.playerHand) })
	g.UI.buttons["end_turn"] = NewButton("End Turn", 300, 300, func() { g.actionEndTurn() })
	// g.deck.moveTopCardToHand(&g.playerHand)
	g.Turn = TurnPlayer
	g.LoadLevelTiles(1)
	g.initPlayerCastle()
	startingEnemies := g.enemyBag.PickStartingEnemies()

	g.spawnSetUpEnemies(startingEnemies)

	if rl.IsMusicStreamPlaying(g.music["iron_at_the_gate"]) {
		rl.StopMusicStream(g.music["iron_at_the_gate"])
	}
	rl.PlayMusicStream(g.music["iron_at_the_gate"])
	g.Round.TurnNumber = 1

}

func (g *Game) actionEndTurn() {
	g.endingTurn = true
}
func (g *Game) spawnSetUpEnemies(enemies []Enemy) {
	for i := range enemies {
		coords := g.GetAllSpawnableTileGridCoords()
		g.PlaceEnemyWithPos(enemies[i], coords[i].X, coords[i].Z)
		// Modify the last added enemy to start falling
		if len(EnemiesInPlay) > 0 {
			idx := len(EnemiesInPlay) - 1
			EnemiesInPlay[idx].isFalling = true
			EnemiesInPlay[idx].visualPos.Y = 20.0
		}
	}
}
func (g *Game) spawnEnemies(enemies []Enemy) {
	for i := range enemies {
		coord := g.GetRandomSpawnableTileGridCoord()
		g.PlaceEnemyWithPos(enemies[i], coord.X, coord.Z)
		// Modify the last added enemy to start falling
		if len(EnemiesInPlay) > 0 {
			idx := len(EnemiesInPlay) - 1
			EnemiesInPlay[idx].isFalling = true
			EnemiesInPlay[idx].visualPos.Y = 20.0
		}
	}
}

func (g *Game) CreateEnemyWave(waveNumber int) {
	if waveNumber == 1 {
		for _, t := range g.levels[g.currentLevel].tiles {
			if t.isSpawn {
				g.NewEnemyWithPos(EnemyTypeKnight, t.gridX, t.gridZ)
			}
		}

	}

}

// func (g *Game) TurnResolve(dt float32) {
// 	g.resolvePlayedCardsrl
// 	if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
// 		g.Turn = TurnComputer
// 		fmt.Println("ENTERING COMPUTER TURN")
// 	}
// }

// func (g *Game) resolvePlayedCards() {
// 	for i := range g.cardsToPlay {
// 		g.cardsToPlay[i].resolve()
// 	}
// 	// g.cardsToPlay = []*Card{}
// }

// func (c *Card) resolve() {
// 	switch c.cardType {
// 	case CardTypeAttackPawn:
// 		for i := range EnemiesInPlay {
// 			if EnemiesInPlay[i].enemyType == EnemyTypePawn {
// 				EnemiesInPlay[i].currentHealth--
// 				if EnemiesInPlay[i].currentHealth <= 0 {
// 					EnemiesInPlay[i] = EnemiesInPlay[len(EnemiesInPlay)-1]
// 					EnemiesInPlay = EnemiesInPlay[:len(EnemiesInPlay)-1]
// 				}
// 				return
// 			}
// 		}
// 	case CardTypeAttackKnight:
// 		for i := range EnemiesInPlay {
// 			if EnemiesInPlay[i].enemyType == EnemyTypeKnight {
// 				EnemiesInPlay[i].currentHealth--
// 				if EnemiesInPlay[i].currentHealth <= 0 {
// 					EnemiesInPlay[i] = EnemiesInPlay[len(EnemiesInPlay)-1]
// 					EnemiesInPlay = EnemiesInPlay[:len(EnemiesInPlay)-1]
// 				}
// 				return
// 			}
// 		}
// 	case CardTypeAttackBishop:
// 		for i := range EnemiesInPlay {
// 			if EnemiesInPlay[i].enemyType == EnemyTypeBishop {
// 				EnemiesInPlay[i].currentHealth--
// 				if EnemiesInPlay[i].currentHealth <= 0 {
// 					EnemiesInPlay[i] = EnemiesInPlay[len(EnemiesInPlay)-1]
// 					EnemiesInPlay = EnemiesInPlay[:len(EnemiesInPlay)-1]
// 				}
// 				return
// 			}
// 		}
// 	}
// }
