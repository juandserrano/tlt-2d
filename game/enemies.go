package game

import rl "github.com/gen2brain/raylib-go/raylib"

type EnemyType int

const (
	EnemyTypePawn EnemyType = iota
	EnemyTypeKnight
	EnemyTypeBishop
	EnemyTypeQueen
	EnemyTypeKing
)

var EnemiesInPlay []Enemy

type Enemy struct {
	model     *rl.Model
	enemyType EnemyType
	posGridX  int
	posGridZ  int
}

func (g *Game) NewEnemy(eType EnemyType, posGridX, posGridZ int) {
	var e Enemy
	e.enemyType = eType
	switch eType {
	case EnemyTypePawn:
		e.model = &g.pawnModel
	}

	e.posGridX = posGridX
	e.posGridZ = posGridZ
	EnemiesInPlay = append(EnemiesInPlay, e)
}

func (e *Enemy) draw(g *Game) {
	pos := g.GetTileCenter(e.posGridX, e.posGridZ)
	rl.DrawModel(*e.model, pos, 1.0, rl.White)
}

func (g *Game) drawEnemies() {
	for i := range EnemiesInPlay {
		EnemiesInPlay[i].draw(g)
	}
}
