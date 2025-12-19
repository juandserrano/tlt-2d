package game

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Hand struct {
	rectangle rl.Rectangle
	cards     []Card
	maxCards  int
}

func (g *Game) NewHand() Hand {
	return Hand{
		rectangle: rl.Rectangle{
			X:      float32(g.Config.Window.Width / 6.0),
			Y:      float32(g.Config.Window.Height) / 6.0 * 4.0,
			Width:  float32(g.Config.Window.Width) / 6.0 * 4,
			Height: float32(g.Config.Window.Height) * 0.3},
		cards:    []Card{},
		maxCards: g.Config.Rules.HandLimit,
	}
}

func (h *Hand) draw() {
	// Draw background box
	rl.DrawRectangleRounded(h.rectangle, 0.2, 0, color.RGBA{50, 50, 50, 50})
}
