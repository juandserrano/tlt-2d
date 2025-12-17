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
	g.CreateEnemyWave(1)

	g.Round.TurnNumber = 1

}

func (g *Game) CreateEnemyWave(waveNumber int) {
	if waveNumber == 1 {
		for _, t := range g.levels[g.currentLevel].tiles {
			if t.isSpawn {
				g.NewEnemy(EnemyTypePawn, t.gridX, t.gridZ)
			}
		}

	}

}
