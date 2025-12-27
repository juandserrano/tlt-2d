package game

import "core:math/rand"
import rl "vendor:raylib"

handlePlayingInput :: proc(g: ^Game, dt: f32) {
	if rl.IsMouseButtonPressed(.LEFT) {
		// 1. UI Layer (Buttons) - Highest Priority
		for _, &b in g.UI.buttons {
			if b.enabled && isMouseOnButton(&b) {
				b.action()
				return // Consume input
			}
		}

		// 2. Hand Layer (Cards)
		// for i := len(g.playerHand.cards) - 1; i >= 0; i -= 1 {
		// Select Card
		// if g.playerHand.cards[i].isMouseOnCard() && g.playerHand.selectedCard == nil {
		// 	rl.SetSoundPitch(g.sounds["card_select"], 0.95 + rand.float32() * 0.1)
		// 	rl.PlaySound(g.sounds["card_select"])
		// 	g.playerHand.cards[i].toggleSelected()
		// 	g.playerHand.selectedCard = g.playerHand.cards[i]
		// 	return // Consume input
		// }
		// Deselect Card
		// 	if g.playerHand.cards[i].isMouseOnCard() && g.playerHand.cards[i].selected {
		// 		rl.SetSoundPitch(g.sounds["card_select"], 0.95 + rand.Float32() * 0.1)
		// 		rl.PlaySound(g.sounds["card_select"])
		// 		g.playerHand.cards[i].toggleSelected()
		// 		g.playerHand.selectedCard = nil
		// 		for e in EnemiesInPlay {
		// 			e.isHighlighted = false
		// 		}
		// 		return // Consume input
		// 	}
		// }

		// 3. Discard Pile
		// for i in g.discardPile.cards {
		// 	if g.discardPile.cards[i].isMouseOnCard() && g.playerHand.selectedCard != nil {
		// 		g.playerHand.moveCardToDiscardPile(g.playerHand.selectedCard, &g.discardPile, g)
		// 		return // Consume input
		// 	}
		// }

		// 4. World Layer (Enemies) - Lowest Priority
		// for &e in EnemiesInPlay {
		// 	if isMouseOnEnemy(g, &e) && e.isHighlighted && g.playerHand.selectedCard != nil {
		// 		if canAttack(g.playerHand.selectedCard, &e) {
		// 			e.isHighlighted = false
		// 			startCardAttack(g, g.playerHand.selectedCard, &e)
		// 			return // Consume input
		// 		}
		// 	}
		// }
	}

}
