package game

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
