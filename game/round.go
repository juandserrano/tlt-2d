package game

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
	g.UI.buttons = append(g.UI.buttons, NewButton("draw", 300, 100, func() { g.deck.drawToTopHand(&g.playerHand) }))
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
