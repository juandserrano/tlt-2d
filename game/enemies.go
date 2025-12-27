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
	model            *rl.Model
	enemyType        EnemyType
	gridPos          GridCoord
	visualPos        rl.Vector3
	moveOnGridX      bool
	maxHealth        int
	currentHealth    int
	attack           int
	healthBarShowing bool
	isHighlighted    bool
	isFalling        bool
	velocityY        float32
}

func (g *Game) NewEnemyWithPos(eType EnemyType, posGridX, posGridZ int) {
	e := g.NewEnemy(eType)

	e.gridPos.X = posGridX
	e.gridPos.Z = posGridZ

	// Initialize visual pos
	worldPos := GridToWorldHex(e.gridPos.X, e.gridPos.Z, HEX_TILE_WIDTH/2.0)
	e.visualPos = rl.Vector3{X: worldPos.X, Y: 0, Z: worldPos.Y}

	EnemiesInPlay = append(EnemiesInPlay, e)
}

func (g *Game) PlaceEnemyWithPos(e Enemy, posGridX, posGridZ int) {
	e.gridPos.X = posGridX
	e.gridPos.Z = posGridZ

	// Initialize visual pos
	worldPos := GridToWorldHex(e.gridPos.X, e.gridPos.Z, HEX_TILE_WIDTH/2.0)
	e.visualPos = rl.Vector3{X: worldPos.X, Y: 0, Z: worldPos.Y}

	EnemiesInPlay = append(EnemiesInPlay, e)
}

func (g *Game) NewEnemy(eType EnemyType) Enemy {
	var e Enemy
	e.enemyType = eType
	e.moveOnGridX = true
	e.healthBarShowing = false
	e.isHighlighted = false
	switch eType {
	case EnemyTypePawn:
		e.model = g.enemyModels[EnemyTypePawn]
		e.maxHealth = g.Config.Enemies.Pawn.Health
		e.attack = g.Config.Enemies.Pawn.Attack
	case EnemyTypeKnight:
		e.model = g.enemyModels[EnemyTypeKnight]
		e.maxHealth = g.Config.Enemies.Knight.Health
		e.attack = g.Config.Enemies.Knight.Attack
	case EnemyTypeBishop:
		e.model = g.enemyModels[EnemyTypeBishop]
		e.maxHealth = g.Config.Enemies.Bishop.Health
		e.attack = g.Config.Enemies.Bishop.Attack
	case EnemyTypeQueen:
		e.model = g.enemyModels[EnemyTypeQueen]
		e.maxHealth = g.Config.Enemies.Bishop.Health
		e.attack = g.Config.Enemies.Bishop.Attack
	case EnemyTypeKing:
		e.model = g.enemyModels[EnemyTypeKing]
		e.maxHealth = g.Config.Enemies.Bishop.Health
		e.attack = g.Config.Enemies.Bishop.Attack
	}
	e.currentHealth = e.maxHealth
	return e

}

func (e *Enemy) Update(dt float32, g *Game) {
	if e.isFalling {
		gravity := float32(50.0)
		e.velocityY -= gravity * dt
		e.visualPos.Y += e.velocityY * dt

		if e.visualPos.Y <= 0 {
			e.visualPos.Y = 0
			e.isFalling = false
			e.velocityY = 0
			// Trigger camera shake
			g.CameraShakeIntensity = 0.5
			rl.SetSoundPitch(g.sounds["falling_impact"], 0.8+rand.Float32()*0.4)
			rl.PlaySound(g.sounds["falling_impact"])
		}
	} else {
		// Target position based on grid
		targetWorld := GridToWorldHex(e.gridPos.X, e.gridPos.Z, HEX_TILE_WIDTH/2.0)
		target := rl.Vector3{X: targetWorld.X, Y: 0, Z: targetWorld.Y}

		// Interpolate
		speed := float32(g.Config.Animations.SlideSpeed)
		e.visualPos = rl.Vector3Lerp(e.visualPos, target, dt*speed)
	}
}

func (e *Enemy) draw(g *Game) {
	if e != nil && e.model != nil {
		// Use visualPos instead of calculating from gridPos
		rl.DrawModelEx(*e.model, e.visualPos, rl.Vector3{X: 0, Y: 1, Z: 0}, float32(util.CalculateRotation(e.visualPos, rl.Vector3{X: 0, Y: 0, Z: 0})), rl.Vector3One(), rl.White)

		// Debug neighbour tile coords
		if g.debugLevel == 2 {
			neighbours := GetNeighbourPositions(e.gridPos)
			for i := range neighbours {
				t := g.GetTileWithGridPos(GridCoord{neighbours[i].X, neighbours[i].Z})
				t.debugDrawGridCoord(rl.Blue)
			}

		}
	}
}

func (g *Game) isMouseOnEnemy(e *Enemy) bool {
	ray := rl.GetScreenToWorldRay(rl.GetMousePosition(), g.camera)

	// Get model bounding box (local space)
	bb := rl.GetModelBoundingBox(*e.model)

	// Get enemy position (world space) - Use visual pos for accurate hit detection during movement?
	// Or stay with grid pos? Let's use visual pos for smoother feel
	pos := e.visualPos

	// Transform bounding box to world space
	bb.Min = rl.Vector3Add(bb.Min, pos)
	bb.Max = rl.Vector3Add(bb.Max, pos)

	rayCollision := rl.GetRayCollisionBox(ray, bb)
	return rayCollision.Hit
}

func (e *Enemy) drawHealthBar(g *Game) {
	barWidth := 50
	barHeight := 10

	// Use visualPos
	targetPos := rl.Vector3{X: e.visualPos.X, Y: 0, Z: e.visualPos.Z}
	screenPosition := rl.GetWorldToScreen(targetPos, g.camera)
	rl.DrawRectangle(int32(screenPosition.X-float32(barWidth)/2.0), int32(screenPosition.Y-float32(barHeight)/2.0), int32(float32(barWidth)*float32(e.currentHealth)/float32(e.maxHealth)), int32(barHeight), rl.Red)

	// rl.PushMatrix()
	// barWidth := 50
	// barHeight := 10
	// enemyWorldPos := GridToWorldHex(e.gridPos.X, e.gridPos.Z, HEX_TILE_WIDTH/2.0)
	// rl.Translatef(enemyWorldPos.X-0.6, 0.1, enemyWorldPos.Y)

	// rl.Rotatef(90, 1, 0, 0)
	// // rl.Rotatef(45, 0, 1, 0)
	// rl.Rotatef(90, 0, 0, 1)
	// rl.Scalef(0.02, 0.02, 0.02)
	// rl.DrawRectangle(-int32(barWidth/2.0), int32(barHeight/2.0), int32(float32(barWidth)*float32(e.currentHealth)/float32(e.maxHealth)), int32(barHeight), rl.Red)
	// rl.PopMatrix()
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

func isTileOccupied(gridPos GridCoord, enemies []Enemy) bool {
	for i := range enemies {
		if enemies[i].currentHealth > 0 && enemies[i].gridPos == gridPos {
			return true
		}
	}
	return false
}

func (e *Enemy) move() {
	neighbourPositions := GetNeighbourPositions(e.gridPos)
	if e.gridPos.X == 0 && e.gridPos.Z == 0 {
		return
	}

	var availablePositions []GridCoord
	for _, pos := range neighbourPositions {
		if !isTileOccupied(pos, EnemiesInPlay) {
			availablePositions = append(availablePositions, pos)
		}
	}

	// If all neighbours are occupied, stay put (or maybe check if we are closer than some available spot?)
	// For now, if blocked, stay put. Ideally we'd check if we are ALREADY at a good spot, but moving towards 0,0 is the goal.
	if len(availablePositions) == 0 {
		return
	}

	closest := closestGridPositionToOrigin(availablePositions)
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

func (g *Game) TurnComputer(dt float32) {
	if g.waitingForSpawnAnimation {
		isAnyFalling := false
		for i := range EnemiesInPlay {
			if EnemiesInPlay[i].isFalling {
				isAnyFalling = true
				break
			}
		}

		if !isAnyFalling && g.CameraShakeIntensity <= 0 {
			g.turnTransitionTimer += dt
			if g.turnTransitionTimer > 1.0 {
				g.turnTransitionTimer = 0
				g.waitingForSpawnAnimation = false
				g.Turn = TurnPlayer
				fmt.Printf("%#v\n", len(EnemiesInPlay))
				// Fade UI back in for player turn
				g.AnimationController.FadeUITo(1.0, 2.0, nil)
			}
		}
		return
	}

	if g.enemyMoveIndex < len(EnemiesInPlay) {
		enemy := &EnemiesInPlay[g.enemyMoveIndex]

		if !g.waitingForMoveAnimation {
			// Start the move
			enemy.move()
			rl.SetSoundPitch(g.sounds["chess_slide"], 0.95+rand.Float32()*0.1)
			rl.PlaySound(g.sounds["chess_slide"])
			g.waitingForMoveAnimation = true
		} else {
			// Check if animation is finished (visual pos close to grid pos)
			targetWorld := GridToWorldHex(enemy.gridPos.X, enemy.gridPos.Z, HEX_TILE_WIDTH/2.0)
			target := rl.Vector3{X: targetWorld.X, Y: 0, Z: targetWorld.Y}

			dist := rl.Vector3Distance(enemy.visualPos, target)

			// Threshold for "arrived"
			if dist < 0.1 {
				g.waitingForMoveAnimation = false
				g.enemyMoveIndex++
			}
		}
		return // Return here to process one enemy per frame sequence
	}

	// All enemies moved
	if len(g.playerHand.cards) < g.playerHand.maxCards {
		g.deck.canDraw = true
	}
	g.spawnEnemies(g.enemyBag.PickRandom(3))

	g.waitingForSpawnAnimation = true
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
