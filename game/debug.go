package game

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) toggleDebug() {
	if rl.IsKeyPressed(rl.KeyG) {
		g.debugLevel++
		if g.debugLevel > 2 {
			g.debugLevel = 0
		}
	}
}
func (g *Game) DrawWorldDebug() {
	// Draw grid coords
	if g.debugLevel == 1 {
		for i := range g.levels[g.currentLevel].tiles {
			g.levels[g.currentLevel].tiles[i].debugDrawGridCoord(rl.Red)
		}

		// Draw ray from mouse to world
		ray := rl.GetScreenToWorldRay(rl.GetMousePosition(), g.camera)
		groundHitInfo := rl.GetRayCollisionQuad(
			ray,
			rl.NewVector3(-100, 0, -100), // Large quad corners
			rl.NewVector3(-100, 0, 100),
			rl.NewVector3(100, 0, 100),
			rl.NewVector3(100, 0, -100),
		)

		if groundHitInfo.Hit {
			fmt.Println("hit")
			// 3. Convert the Hit Point (3D World) to Hex Grid (2D Logical)
			// Note: We map World X -> Hex X, and World Z -> Hex Y
			rl.DrawRay(ray, rl.Green)
		} else {
			fmt.Println("no hit")
		}

	}

}

func (g *Game) DrawStaticDebug() {
	rl.DrawFPS(10, 10)
	rl.DrawText("DEBUG MODE", int32(g.Config.Window.Width)-100, 10, 16, rl.Red)
	rl.DrawText(fmt.Sprintf("mousepos: %v", rl.GetMousePosition()), 100, 100, 20, rl.Blue)
}

func (t *Tile) debugDrawGridCoord(color color.RGBA) {
	// --- DRAW TEXT 3D ---
	// 1. Push the current matrix so we don't mess up other 3D objects
	rl.PushMatrix()

	// 2. Move to the position in 3D space (X, Y, Z)
	rl.Translatef(t.position.X, 0.1, t.position.Z)

	// 3. Rotate the text.
	// By default, text lies flat on the floor looking up.
	// Rotate 90 degrees on X to make it stand up.
	// Rotate 180 degrees on Y because text usually renders "backwards" in 3D look-at logic.
	rl.Rotatef(90, 1, 0, 0)
	// rl.Rotatef(90, 0, 1, 0)
	rl.Rotatef(90, 0, 0, 1)

	// 4. Scale it DOWN.
	// Standard font size 20 is "20 meters" high in 3D.
	// We scale by 0.1 to make it manageable.
	rl.Scalef(0.01, 0.01, 0.01)

	// 5. Draw the text (Relative to 0,0 because we already translated the matrix)
	// We center the text by calculating width/2
	text := fmt.Sprintf("(%d, %d)", t.gridX, t.gridZ)
	fontSize := int32(20)
	textWidth := rl.MeasureText(text, fontSize)

	rl.DrawText(text, -textWidth/2, 0, fontSize, color)

	// 6. Restore the matrix
	rl.PopMatrix()
	// --------------------

}
