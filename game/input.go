package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) handlePlayingInput(dt float32) {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		for i := len(g.playerHand.cards) - 1; i >= 0; i-- {
			if g.playerHand.cards[i].isMouseOnCard() && g.playerHand.selectedCard == nil {
				g.playerHand.cards[i].toggleSelected()
				g.playerHand.selectedCard = &g.playerHand.cards[i]
				break
			}
			if g.playerHand.cards[i].isMouseOnCard() && g.playerHand.cards[i].selected {
				g.playerHand.cards[i].toggleSelected()
				g.playerHand.selectedCard = nil
				for i := range EnemiesInPlay {
					EnemiesInPlay[i].isHighlighted = false
				}
				break
			}
		}
		for i := range EnemiesInPlay {
			if g.isMouseOnEnemy(&EnemiesInPlay[i]) && EnemiesInPlay[i].isHighlighted && g.playerHand.selectedCard != nil {
				EnemiesInPlay[i].isHighlighted = false

				g.playerHand.selectedCard.attackEnemy(&EnemiesInPlay[i], &g.playerHand)

			}
		}

		for _, v := range g.UI.buttons {
			if v.enabled && v.MouseOnButton() {
				v.action()
			}

		}
	}

}
