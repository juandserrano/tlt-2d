package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	g.deck = g.NewDeck()
	g.UI.buttons["draw"] = NewButton("draw", 300, 100, func() { g.drawToTopHand(&g.playerHand) })
	g.UI.buttons["play"] = NewButton("play", 300, 300, func() { g.playerHand.playSelected(g) })
	// g.deck.moveTopCardToHand(&g.playerHand)
	g.Turn = TurnPlayer
	g.LoadLevelTiles(1)
	g.initPlayerCastle()
	startingEnemies := g.enemyBag.PickStartingEnemies()

	g.spawnEnemies(startingEnemies)

	// g.CreateEnemyWave(1)

	g.Round.TurnNumber = 1

}

func (g *Game) spawnEnemies(enemies []Enemy) {
	for _, t := range g.levels[g.currentLevel].tiles {
		if t.isSpawn {
			g.PlaceEnemyWithPos(enemies[0], t.gridX, t.gridZ)
			enemies[0] = enemies[len(enemies)-1]
			enemies = enemies[:len(enemies)-1]
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

func (g *Game) TurnResolve(dt float32) {
	g.resolvePlayedCards()
	if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		g.Turn = TurnComputer
		fmt.Println("ENTERING COMPUTER TURN")
	}
}

func (g *Game) resolvePlayedCards() {
	for i := range g.cardsToPlay {
		g.cardsToPlay[i].resolve()
	}
	// g.cardsToPlay = []*Card{}
}

func (c *Card) resolve() {
	switch c.cardType {
	case CardTypeAttackPawn:
		fmt.Println("Its a pwn cards")
		for i := range EnemiesInPlay {
			if EnemiesInPlay[i].enemyType == EnemyTypePawn {
				fmt.Printf("Resolving:\n%#v\n", c)
				fmt.Println(EnemiesInPlay[i].currentHealth)
				EnemiesInPlay[i].currentHealth--
				fmt.Println(EnemiesInPlay[i].currentHealth)
				return
			}
		}
	case CardTypeAttackKnight:
		for i := range EnemiesInPlay {
			if EnemiesInPlay[i].enemyType == EnemyTypeKnight {
				EnemiesInPlay[i].currentHealth--
				break
			}
		}
	case CardTypeAttackBishop:
		for i := range EnemiesInPlay {
			if EnemiesInPlay[i].enemyType == EnemyTypeBishop {
				EnemiesInPlay[i].currentHealth--
				break
			}
		}
	}
}
