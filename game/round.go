package game

import "fmt"

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
	fmt.Println(len(g.enemyBag.enemies))
	g.Turn = TurnPlayer
	g.LoadLevelTiles(1)
	g.initPlayerCastle()
	startingEnemies := g.enemyBag.PickStartingEnemies()
	fmt.Println(len(g.enemyBag.enemies))

	g.spawnEnemies(startingEnemies)

	// g.CreateEnemyWave(1)

	g.Round.TurnNumber = 1

}

func (g *Game) spawnEnemies(enemies []Enemy) {
	for _, t := range g.levels[g.currentLevel].tiles {
		if t.isSpawn {
			g.PlaceEnemyWithPos(enemies[0], t.gridX, t.gridZ)
			fmt.Println("Spawning enemy type", enemies[0].enemyType)
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
