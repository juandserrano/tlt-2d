package game

import "core:fmt"
import "core:strings"
import rl "vendor:raylib"

toggleDebug :: proc(g: ^Game) {
	if rl.IsKeyPressed(.G) {
		g.debugLevel += 1
		if g.debugLevel > 2 {
			g.debugLevel = 0
		}
	}
}

drawWorldDebug :: proc(g: ^Game) {
	// Draw grid coords
	if g.debugLevel == 1 {
		for &t in g.levels[g.currentLevel].tiles {
			// debugDrawGridCoord(&t, rl.Red)
		}
	}

}

DrawStaticDebug :: proc(g: ^Game) {
	rl.DrawFPS(10, 10)
	rl.DrawText("DEBUG MODE", i32(g.Config.Window.Width) - 100, 10, 16, rl.RED)
	text := fmt.tprintf("mousepos: %v", rl.GetMousePosition())
	c_str := strings.clone_to_cstring(text)
	defer delete(c_str)
	rl.DrawText(c_str, 100, 100, 20, rl.BLUE)
}

terminalDebug :: proc(g: ^Game) {
	if g.frameCount % 30 == 0 {
		for e in EnemiesInPlay {
			if e.enemyType == .EnemyTypeBishop {
			}
		}

	}
}

// debugDrawGridCoord :: proc(t: ^Tile, color: rl.Color) {
// 	// --- DRAW TEXT 3D ---
// 	// 1. Push the current matrix so we don't mess up other 3D objects
// 	rl.PushMatrix()

// 	// 2. Move to the position in 3D space (X, Y, Z)
// 	rl.Translatef(t.position.x, 0.1, t.position.z)

// 	// 3. Rotate the text.
// 	// By default, text lies flat on the floor looking up.
// 	// Rotate 90 degrees on X to make it stand up.
// 	// Rotate 180 degrees on Y because text usually renders "backwards" in 3D look-at logic.
// 	rl.Rotatef(90, 1, 0, 0)
// 	// rl.Rotatef(90, 0, 1, 0)
// 	rl.Rotatef(90, 0, 0, 1)

// 	// 4. Scale it DOWN.
// 	// Standard font size 20 is "20 meters" high in 3D.
// 	// We scale by 0.1 to make it manageable.
// 	rl.Scalef(0.01, 0.01, 0.01)

// 	// 5. Draw the text (Relative to 0,0 because we already translated the matrix)
// 	// We center the text by calculating width/2
// 	text := fmt.Sprintf("(%d, %d)", t.gridX, t.gridZ)
// 	fontSize := int32(20)
// 	textWidth := rl.MeasureText(text, fontSize)

// 	rl.DrawText(text, -textWidth / 2, 0, fontSize, color)

// 	// 6. Restore the matrix
// 	rl.PopMatrix()
// 	// --------------------

// }
