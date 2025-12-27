package game

import "core:fmt"
import "core:math"
import "core:math/rand"
import "util"
import rl "vendor:raylib"

EnemyBag :: struct {
	enemies: []Enemy,
}

EnemyType :: enum {
	EnemyTypePawn,
	EnemyTypeKnight,
	EnemyTypeBishop,
	EnemyTypeQueen,
	EnemyTypeKing,
}

EnemiesInPlay: []Enemy

Enemy :: struct {
	model:            ^rl.Model,
	enemyType:        EnemyType,
	gridPos:          GridCoord,
	visualPos:        rl.Vector3,
	moveOnGridX:      bool,
	maxHealth:        int,
	currentHealth:    int,
	attack:           int,
	healthBarShowing: bool,
	isHighlighted:    bool,
	isFalling:        bool,
	velocityY:        f32,
}

drawEnemyHealthBar :: proc(e: ^Enemy, g: ^Game) {
	barWidth := 50
	barHeight := 10

	// Use visualPos
	targetPos := rl.Vector3{e.visualPos.x, 0, e.visualPos.z}
	screenPosition := rl.GetWorldToScreen(targetPos, g.camera)
	rl.DrawRectangle(
		i32(screenPosition.x - f32(barWidth) / 2.0),
		i32(screenPosition.y - f32(barHeight) / 2.0),
		i32(f32(barWidth) * f32(e.currentHealth) / f32(e.maxHealth)),
		i32(barHeight),
		rl.RED,
	)
}

drawEnemy :: proc(e: ^Enemy, g: ^Game) {
	if (e != nil && e.model != nil) {
		// Use visualPos instead of calculating from gridPos
		rotation := util.CalculateRotation(e.visualPos, rl.Vector3{0, 0, 0})
		rl.DrawModelEx(
			e.model^,
			e.visualPos,
			rl.Vector3{0, 1, 0},
			f32(rotation),
			rl.Vector3(1),
			rl.WHITE,
		)

		// Debug neighbour tile coords
		if g.debugLevel == 2 {
			neighbours: []GridCoord
			GetNeighbourPositions(e.gridPos, &neighbours)
			for n in neighbours {
				// t := GetTileWithGridPos(g, GridCoord{n.x, n.z})
				// debugDrawGridCoord(t, rl.BLUE)
			}

		}
	}
}

drawEnemies :: proc(g: ^Game) {
	for &e in EnemiesInPlay {
		drawEnemy(&e, g)
	}
}

UpdateEnemies :: proc() {
	for &e in EnemiesInPlay {
		moveEnemy(&e)
	}
}

closestGridPositionToOrigin :: proc(gridPositions: [dynamic]GridCoord) -> GridCoord {
	closestDistance := 99999.0
	closestGridCoord: GridCoord
	for i in gridPositions {
		pos := GridToWorldHex(i.x, i.z, HEX_TILE_WIDTH / 2.0)
		d := math.sqrt(math.pow(f64(pos.x), 2) + math.pow(f64(pos.y), 2))
		if d < f64(closestDistance) {
			closestGridCoord = i
			closestDistance = d
		}
	}
	return closestGridCoord
}

isTileOccupied :: proc(gridPos: GridCoord, enemies: []Enemy) -> bool {
	for e in enemies {
		if e.currentHealth > 0 && e.gridPos == gridPos {
			return true
		}
	}
	return false
}

moveEnemy :: proc(e: ^Enemy) {
	neighbourPositions: []GridCoord
	GetNeighbourPositions(e.gridPos, &neighbourPositions)
	if e.gridPos.x == 0 && e.gridPos.z == 0 {
		return
	}

	availablePositions: [dynamic]GridCoord
	for pos in neighbourPositions {
		if !isTileOccupied(pos, EnemiesInPlay) {
			append(&availablePositions, pos)
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

GetNeighbourPositions :: proc(c: GridCoord, result: ^[]GridCoord) {
	if c.x % 2 != 0 {
		result^ = {
			{c.x - 1, c.z}, //
			{c.x - 1, c.z + 1}, //
			{c.x, c.z + 1}, //
			{c.x, c.z - 1}, //
			{c.x + 1, c.z + 1}, //
			{c.x + 1, c.z}, //
		}
		return
	}
	result^ = {
		{c.x - 1, c.z}, //
		{c.x - 1, c.z - 1}, //
		{c.x, c.z + 1}, //
		{c.x, c.z - 1}, //
		{c.x + 1, c.z - 1}, //
		{c.x + 1, c.z}, //
	}

}

TurnComputer :: proc(g: ^Game, dt: f32) {
	if g.waitingForSpawnAnimation {
		isAnyFalling := false
		for e in EnemiesInPlay {
			if e.isFalling {
				isAnyFalling = true
				break
			}
		}

		if !isAnyFalling && g.CameraShakeIntensity <= 0 {
			g.turnTransitionTimer += dt
			if g.turnTransitionTimer > 1.0 {
				g.turnTransitionTimer = 0
				g.waitingForSpawnAnimation = false
				g.Turn = .TurnPlayer
				fmt.printf("%#v\n", len(EnemiesInPlay))
				// Fade UI back in for player turn
				// g.AnimationController.FadeUITo(1.0, 2.0, nil)
			}
		}
		return
	}

	if g.enemyMoveIndex < len(EnemiesInPlay) {
		enemy := &EnemiesInPlay[g.enemyMoveIndex]

		if !g.waitingForMoveAnimation {
			// Start the move
			moveEnemy(enemy)
			rl.SetSoundPitch(g.sounds["chess_slide"], 0.95 + rand.float32() * 0.1)
			rl.PlaySound(g.sounds["chess_slide"])
			g.waitingForMoveAnimation = true
		} else {
			// Check if animation is finished (visual pos close to grid pos)
			targetWorld := GridToWorldHex(enemy.gridPos.x, enemy.gridPos.z, HEX_TILE_WIDTH / 2.0)
			target := rl.Vector3{targetWorld.x, 0, targetWorld.y}

			dist := rl.Vector3Distance(enemy.visualPos, target)

			// Threshold for "arrived"
			if dist < 0.1 {
				g.waitingForMoveAnimation = false
				g.enemyMoveIndex += 1
			}
		}
		return // Return here to process one enemy per frame sequence
	}

	// All enemies moved
	// if len(g.playerHand.cards) < g.playerHand.maxCards {
	// 	g.deck.canDraw = true
	// }
	// g.spawnEnemies(g.enemyBag.PickRandom(3))

	// g.waitingForSpawnAnimation = true
}

NewEnemyBag :: proc(g: ^Game) -> EnemyBag {
	enemies: []Enemy
	pawnQty := 20
	knightQty := 10
	bishopQty := 5

	// for range pawnQty {
	// 	enemies = append(enemies, g.NewEnemy(EnemyTypePawn))
	// }
	// for range knightQty {
	// 	enemies = append(enemies, g.NewEnemy(EnemyTypeKnight))
	// }
	// for range bishopQty {
	// 	enemies = append(enemies, g.NewEnemy(EnemyTypeBishop))
	// }
	// enemies = append(enemies, g.NewEnemy(EnemyTypeQueen))
	// enemies = append(enemies, g.NewEnemy(EnemyTypeKing))

	bag := EnemyBag {
		enemies = enemies,
	}

	return bag

}

PickRandomEnemy :: proc(b: ^EnemyBag, qty: int) -> [dynamic]Enemy {
	picked: [dynamic]Enemy
	for i in 0 ..= qty {
		if len(b.enemies) == 0 {
			fmt.println("No more enemies in the bag")
			break
		}
		idx := rand.int_range(0, len(b.enemies))
		append(&picked, b.enemies[idx])
		b.enemies[idx] = b.enemies[len(b.enemies) - 1]
		b.enemies = b.enemies[:len(b.enemies) - 1]
	}
	return picked
}

ErrorType :: enum {
	ErrorNoMoreEnemies,
}
PickOneFromType :: proc(b: ^EnemyBag, eType: EnemyType) -> (Enemy, ErrorType) {
	picked: Enemy
	if len(b.enemies) == 0 {
		return picked, .ErrorNoMoreEnemies
	}
	for &e in b.enemies {
		if e.enemyType == eType {
			picked = e
			e = b.enemies[len(b.enemies) - 1]
			b.enemies = b.enemies[:len(b.enemies) - 1]
			return picked, nil
		}
	}
	return picked, .ErrorNoMoreEnemies
}

PickStartingEnemies :: proc(b: ^EnemyBag) -> [dynamic]Enemy {
	startingEnemies: [dynamic]Enemy
	pawnQty := 3
	knightQty := 2
	bishopQty := 1
	for i in 0 ..= pawnQty {
		pawn, err := PickOneFromType(b, .EnemyTypePawn)
		if err != nil {
			return nil
		}
		append(&startingEnemies, pawn)
	}
	for i in 0 ..= knightQty {
		knight, err := PickOneFromType(b, .EnemyTypeKnight)
		if err != nil {
			return nil
		}
		append(&startingEnemies, knight)
	}
	for i in 0 ..= bishopQty {
		bishop, err := PickOneFromType(b, .EnemyTypeBishop)
		if err != nil {
			return nil
		}
		append(&startingEnemies, bishop)
	}
	return startingEnemies

}
