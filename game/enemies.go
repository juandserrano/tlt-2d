package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	model       *rl.Model
	enemyType   EnemyType
	gridPos     GridCoord
	moveOnGridX bool
}

func (g *Game) NewEnemy(eType EnemyType, posGridX, posGridZ int) {
	var e Enemy
	e.enemyType = eType
	e.moveOnGridX = true
	switch eType {
	case EnemyTypePawn:
		e.model = &g.pawnModel
	}

	e.gridPos.X = posGridX
	e.gridPos.Z = posGridZ
	EnemiesInPlay = append(EnemiesInPlay, e)
}

func (e *Enemy) draw(g *Game) {
	pos := g.GetTileCenter(e.gridPos)
	rl.DrawModel(*e.model, pos, 1.0, rl.White)
}

func (g *Game) drawEnemies() {
	for i := range EnemiesInPlay {
		EnemiesInPlay[i].draw(g)
	}
}

func (g *Game) UpdateEnemies() {
	for i := range EnemiesInPlay {
		EnemiesInPlay[i].move()
	}
}

func (e *Enemy) move() {
	neighbourPositions := GetNeighbourPositions(e.gridPos)
	fmt.Printf("%v", neighbourPositions)
	switch e.enemyType {
	case EnemyTypePawn:
		if e.moveOnGridX {
			e.gridPos.X--
		} else {
			e.gridPos.Z--
		}
		e.moveOnGridX = !e.moveOnGridX

	}
}

func GetNeighbourPositions(c GridCoord) []GridCoord {
	return []GridCoord{
		{X: c.X - 1, Z: c.Z},
		{X: c.X - 1, Z: c.Z - 1},
		{X: c.X - 1, Z: c.Z + 1},
		{X: c.X, Z: c.Z + 1},
		{X: c.X, Z: c.Z - 1},
		{X: c.X + 1, Z: c.Z + 1},
		{X: c.X + 1, Z: c.Z},
		{X: c.X + 1, Z: c.Z - 1},
	}

}
