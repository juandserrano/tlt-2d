package game

import (
	"fmt"
	"math"

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
	if g.debugLevel == 2 {
		neighbours := GetNeighbourPositions(e.gridPos)
		for i := range neighbours {
			t := g.GetTileWithGridPos(GridCoord{neighbours[i].X, neighbours[i].Z})
			t.debugDrawGridCoord(rl.Blue)
		}

	}
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

func closestGridPositionToOrigin(gridPositions []GridCoord) GridCoord {
	closestDistance := 99999.0
	var closestGridCoord GridCoord
	for i := range gridPositions {
		pos := GridToWorldHex(gridPositions[i].X, gridPositions[i].Z, HEX_TILE_WIDTH/2.0)
		d := math.Sqrt(math.Pow(float64(pos.X), 2) + math.Pow(float64(pos.Y), 2))
		if d < float64(closestDistance) {
			closestGridCoord = gridPositions[i]
			closestDistance = d
		}
	}
	return closestGridCoord
}

func (e *Enemy) move() {
	neighbourPositions := GetNeighbourPositions(e.gridPos)
	closest := closestGridPositionToOrigin(neighbourPositions)
	if e.gridPos.X == 0 && e.gridPos.Z == 0 {
		closest = e.gridPos
	}
	switch e.enemyType {
	case EnemyTypePawn:
		e.gridPos = closest
	}
}

func GetNeighbourPositions(c GridCoord) []GridCoord {
	if c.X%2 != 0 {
		return []GridCoord{
			{X: c.X - 1, Z: c.Z},     //
			{X: c.X - 1, Z: c.Z + 1}, //
			{X: c.X, Z: c.Z + 1},     //
			{X: c.X, Z: c.Z - 1},     //
			{X: c.X + 1, Z: c.Z + 1}, //
			{X: c.X + 1, Z: c.Z},     //
		}
	}
	return []GridCoord{
		{X: c.X - 1, Z: c.Z},     //
		{X: c.X - 1, Z: c.Z - 1}, //
		{X: c.X, Z: c.Z + 1},     //
		{X: c.X, Z: c.Z - 1},     //
		{X: c.X + 1, Z: c.Z - 1}, //
		{X: c.X + 1, Z: c.Z},     //
	}

}

func (g *Game) TurnResolve(dt float32) {
	fmt.Println("Resolving...")
	if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		g.Turn = TurnComputer
		fmt.Println("ENTERING COMPUTER TURN")
	}
}

func (g *Game) TurnComputer(dt float32) {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		g.NextTurn()
	}
	if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		g.Turn = TurnPlayer
		fmt.Println("ENTERING PLAYER TURN")
	}

}
