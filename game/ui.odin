package game

import "core:strings"
import rl "vendor:raylib"

UI :: struct {
	buttons: map[string]Button,
}

Button :: struct {
	text:     string,
	position: rl.Vector2,
	rX:       int,
	rY:       int,
	action:   proc(),
	enabled:  bool,
}

drawUI :: proc(g: ^Game) {
	for &e in EnemiesInPlay {
		if e.healthBarShowing {
			drawEnemyHealthBar(&e, g)
		}

	}
	for _, &b in g.UI.buttons {
		if b.enabled {
			drawButton(&b)
		}
	}
}

NewButton :: proc(text: string, positionX, positionY: int, action: proc()) -> Button {
	return (Button {
				text = text,
				position = rl.Vector2{f32(positionX), f32(positionY)},
				rX = 40,
				rY = 20,
				action = action,
				enabled = false,
			})
}

MouseOnButton :: proc(b: ^Button) -> bool {
	point := rl.GetMousePosition()
	// Avoid division by zero
	if b.rX <= 0 || b.rY <= 0 {
		return false
	}

	// Normalized squared distance from center
	dx := point.x - b.position.x
	dy := point.y - b.position.y

	termX := (dx * dx) / f32(b.rX * b.rX)
	termY := (dy * dy) / f32(b.rY * b.rY)

	// If sum is <= 1.0, point is inside or on edge
	return (termX + termY) <= 1.0
}

drawButton :: proc(b: ^Button) {
	rl.DrawEllipse(i32(b.position.x), i32(b.position.y), f32(b.rX), f32(b.rY), rl.BLACK)
	c_str := strings.clone_to_cstring(b.text)
	defer delete(c_str)
	textWidth := rl.MeasureText(c_str, 15)
	rl.DrawText(c_str, i32(b.position.x) - textWidth / 2, i32(b.position.y) - 7, 15, rl.BLUE)
}
