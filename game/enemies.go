package game

import (
	"fmt"
	"juandserrano/tlt-2d/game/util"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemyBag struct {
	enemies []Enemy
}

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

func (g *Game) NewEnemyWithPos(eType EnemyType, posGridX, posGridZ int) {
	e := g.NewEnemy(eType)

	e.gridPos.X = posGridX
	e.gridPos.Z = posGridZ

	EnemiesInPlay = append(EnemiesInPlay, e)
}

func (g *Game) PlaceEnemyWithPos(e Enemy, posGridX, posGridZ int) {
	e.gridPos.X = posGridX
	e.gridPos.Z = posGridZ
	EnemiesInPlay = append(EnemiesInPlay, e)
}

func (g *Game) NewEnemy(eType EnemyType) Enemy {
	var e Enemy
	e.enemyType = eType
	e.moveOnGridX = true
	switch eType {
	case EnemyTypePawn:
		e.model = g.enemyModels[EnemyTypePawn]
	case EnemyTypeKnight:
		e.model = g.enemyModels[EnemyTypeKnight]
	case EnemyTypeBishop:
		e.model = g.enemyModels[EnemyTypeBishop]
	}
	return e

}

func (e *Enemy) draw(g *Game) {
	pos := g.GetTileCenter(e.gridPos)
	rl.DrawModelEx(*e.model, pos, rl.Vector3{0, 1, 0}, float32(util.CalculateRotation(pos, rl.Vector3{0, 0, 0})), rl.Vector3One(), rl.White)
	// Debug neighbour tile coords
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
	e.gridPos = closest
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

func (g *Game) NewEnemyBag() EnemyBag {
	var enemies []Enemy
	pawnQty := 20
	knightQty := 10
	bishopQty := 5

	for range pawnQty {
		enemies = append(enemies, g.NewEnemy(EnemyTypePawn))
	}
	for range knightQty {
		enemies = append(enemies, g.NewEnemy(EnemyTypeKnight))
	}
	for range bishopQty {
		enemies = append(enemies, g.NewEnemy(EnemyTypeBishop))
	}
	enemies = append(enemies, g.NewEnemy(EnemyTypeQueen))
	enemies = append(enemies, g.NewEnemy(EnemyTypeKing))

	bag := EnemyBag{
		enemies: enemies,
	}

	return bag

}

func (b *EnemyBag) PickRandom(qty int) []Enemy {
	var picked []Enemy
	for range qty {
		if len(b.enemies) == 0 {
			fmt.Println("No more enemies in the bag")
			break
		}
		idx := rand.Intn(len(b.enemies))
		picked = append(picked, b.enemies[idx])
		b.enemies[idx] = b.enemies[len(b.enemies)-1]
		b.enemies = b.enemies[:len(b.enemies)-1]
	}
	return picked
}

func (b *EnemyBag) PickOneFromType(eType EnemyType) (Enemy, error) {
	var picked Enemy
	if len(b.enemies) == 0 {
		return picked, fmt.Errorf("no more enemies in the bag")
	}
	for i := range b.enemies {
		if b.enemies[i].enemyType == eType {
			picked = b.enemies[i]
			b.enemies[i] = b.enemies[len(b.enemies)-1]
			b.enemies = b.enemies[:len(b.enemies)-1]
			return picked, nil
		}
	}
	return picked, fmt.Errorf("no more enemies of type %v", eType)
}

func (b *EnemyBag) PickStartingEnemies() []Enemy {
	var startingEnemies []Enemy
	pawnQty := 3
	knightQty := 2
	bishopQty := 1
	for range pawnQty {
		pawn, err := b.PickOneFromType(EnemyTypePawn)
		if err != nil {
			return nil
		}
		startingEnemies = append(startingEnemies, pawn)
	}
	for range knightQty {
		knight, err := b.PickOneFromType(EnemyTypeKnight)
		if err != nil {
			return nil
		}
		startingEnemies = append(startingEnemies, knight)
	}
	for range bishopQty {
		bishop, err := b.PickOneFromType(EnemyTypeBishop)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		startingEnemies = append(startingEnemies, bishop)
	}
	return startingEnemies

}
