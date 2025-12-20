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

	// Draw cards
	for _, c := range h.cards {
		c.draw()
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
