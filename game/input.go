package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) handlePlayingInput(dt float32) {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		for i := range g.playerHand.cards {
			if g.playerHand.cards[i].isMouseOnCard() {
				g.playerHand.cards[i].toggleSelected()
			}

		}
		for i := range g.UI.buttons {
			if g.UI.buttons[i].MouseOnButton() {
				g.UI.buttons[i].action()
			}

		}
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonRight) && rl.IsKeyDown(rl.KeyLeftShift) {
		g.Turn = TurnResolving
		fmt.Println("ENTERING RESOLVING STATE")
	}

	// if rl.IsKeyPressed(rl.KeyA) && rl.IsKeyDown(rl.KeyLeftShift) {
	// 	g.State = StateWorldEditor
	// 	fmt.Println("ENTERING WORLD EDITOR")
	// }
}
