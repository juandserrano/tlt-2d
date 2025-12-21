package game

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type HandPosition struct {
	available bool
	position  rl.Vector2
}
type Hand struct {
	rectangle     rl.Rectangle
	cards         []Card
	maxCards      int
	cardPositions []HandPosition
	// indexesPlayed  []int
	selectedCard *Card
}

func (g *Game) NewHand() Hand {
	h := Hand{
		rectangle: rl.Rectangle{
			X:      float32(g.Config.Window.Width / 6.0),
			Y:      float32(g.Config.Window.Height) / 6.0 * 4.0,
			Width:  float32(g.Config.Window.Width) / 6.0 * 4,
			Height: float32(g.Config.Window.Height) * 0.3},
		cards:    []Card{},
		maxCards: g.Config.Rules.HandLimit,
	}

	for i := range h.maxCards {
		h.cardPositions = append(h.cardPositions, HandPosition{
			available: true,
			position: rl.Vector2{
				X: h.rectangle.X + (float32(i+1) * h.rectangle.Width / float32(h.maxCards)),
				Y: h.rectangle.Y + h.rectangle.Height/2.0,
			},
		})
	}

	return h
}

func (h *Hand) draw() {
	// Draw background box
	rl.DrawRectangleRounded(h.rectangle, 0.2, 0, color.RGBA{50, 50, 50, 50})

	// Draw cards by hand position
	for i := range h.cardPositions {
		for _, c := range h.cards {
			if c.positionInHand == i {
				c.draw()
			}
		}
	}
}

func (h *Hand) nextAvailablePosition() (int, rl.Vector2, error) {
	for i := range h.cardPositions {
		if h.cardPositions[i].available {
			return i, h.cardPositions[i].position, nil
		}
	}
	return 999, rl.Vector2Zero(), fmt.Errorf("hand is full")
}

// func (h *Hand) playSelected(g *Game) {
// 	g.cardsToPlay = []*Card{}
// 	for i := range h.cards {
// 		if h.cards[i].selected {
// 			h.cards[i].addToplay(g)
// 			h.cardPositions[h.cards[i].positionInHand].available = true
// 			h.indexesPlayed = append(h.indexesPlayed, i)
// 		}
// 	}
// 	h.moveCardsToDiscardPile(h.indexesPlayed, g)
// 	h.indexesPlayed = []int{}
// 	g.Turn = TurnResolving

// }

func (h *Hand) UpdateHand() []Card {
	n := 0
	for _, c := range h.cards {
		if !c.selected {
			h.cards[n] = c
			n++
		}
	}
	return h.cards[:n]
}

func (h *Hand) moveCardToDiscardPile(c *Card) {
	h.selectedCard = nil
	h.cardPositions[c.positionInHand].available = true
	for i := range h.cards {
		if h.cards[i].id == c.id {
			h.cards[i] = h.cards[len(h.cards)-1]
			h.cards = h.cards[:len(h.cards)-1]
			break
		}
	}
	// h.cards = h.UpdateHand()
}

func (g *Game) reorderHand() {
	for i := range g.playerHand.cardPositions {
		g.playerHand.cardPositions[i].available = true
	}

	for i := range g.playerHand.cards {
		g.playerHand.cards[i].positionInHand = i
		g.playerHand.cardPositions[i].available = false
	}
}
