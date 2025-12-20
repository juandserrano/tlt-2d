package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) handlePlayingInput(dt float32) {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		for i := len(g.playerHand.cards) - 1; i >= 0; i-- {
			if g.playerHand.cards[i].isMouseOnCard() {
				g.playerHand.cards[i].toggleSelected()
				break
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
