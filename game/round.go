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
	g.Turn = TurnPlayer
	g.LoadLevelTiles(1)
	g.initPlayerCastle()
	g.CreateStartEnemies()

	g.Round.TurnNumber = 1

}

func (g *Game) CreateStartEnemies() {
	g.NewEnemy(EnemyTypePawn, 5, 5)
	g.testPawn = EnemiesInPlay[0]

}
