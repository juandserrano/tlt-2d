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
	cards         []*Card
	maxCards      int
	cardPositions []HandPosition
	// indexesPlayed  []int
	selectedCard *Card
}

func (g *Game) OnWindowSizeUpdate() {
	screenMid := rl.GetScreenWidth() / 2
	rectWidth := 600
	x := screenMid - rectWidth/2

	g.playerHand.rectangle.X = float32(x)
	g.playerHand.rectangle.Width = float32(rectWidth)
	g.playerHand.rectangle.Y = float32(rl.GetScreenHeight()) - 30 - g.playerHand.rectangle.Height

	// update Card positions
	for i := range g.playerHand.cardPositions {
		g.playerHand.cardPositions[i].position.X = g.playerHand.rectangle.X + (float32(i+1) * 100) // g.playerHand.rectangle.Width / float32(g.playerHand.maxCards))
		g.playerHand.cardPositions[i].position.Y = g.playerHand.rectangle.Y + g.playerHand.rectangle.Height/2.0
	}
	for i := range g.playerHand.cards {
		posIndex := g.playerHand.cards[i].positionInHand
		g.playerHand.cards[i].position = rl.Vector2Add(g.playerHand.cardPositions[posIndex].position, rl.Vector2{X: float32(-g.cardTextures[CardTypeAttackPawn].Width / 2), Y: float32(-g.cardTextures[CardTypeAttackPawn].Height / 2)})
	}

	g.discardPile.position.X = float32(rl.GetScreenWidth()) - float32(g.cardTextures[CardTypeAttackPawn].Width) - 10

	for i := range g.discardPile.cards {
		g.discardPile.cards[i].position = rl.Vector2Add(g.discardPile.position, rl.Vector2{X: float32(i) * 3, Y: 0})
	}

}

func (g *Game) NewHand() Hand {
	x := float32(g.Config.Window.Width / 6.0)
	xEnd := x * 5
	rectWidth := xEnd - x
	h := Hand{
		rectangle: rl.Rectangle{
			X:      x,
			Y:      float32(g.Config.Window.Height) / 6.0 * 4.0,
			Width:  rectWidth,
			Height: float32(g.Config.Window.Height) * 0.3},
		cards:    []*Card{},
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

func (h *Hand) draw(alpha float32) {
	// Draw background box
	// alpha255 := uint8(alpha * 255) // Unused variable
	rl.DrawRectangleRounded(h.rectangle, 0.2, 0, color.RGBA{50, 50, 50, uint8(float32(50) * alpha)})

	// Draw cards by hand position
	for i := range h.cardPositions {
		for _, c := range h.cards {
			if c.positionInHand == i {
				c.draw(1, alpha)
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

func (h *Hand) UpdateHand() []*Card {
	n := 0
	for _, c := range h.cards {
		if !c.selected {
			h.cards[n] = c
			n++
		}
	}
	return h.cards[:n]
}

func (h *Hand) moveCardToDiscardPile(c *Card, discardPile *Deck) {
	h.selectedCard = nil
	h.cardPositions[c.positionInHand].available = true
	c.isShowing = true
	c.selected = false
	c.position = discardPile.position
	discardPile.Push(c)
	for i := range h.cards {
		if h.cards[i].id == c.id {
			h.cards[i] = h.cards[len(h.cards)-1]
			h.cards = h.cards[:len(h.cards)-1]
			break
		}
	}
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
