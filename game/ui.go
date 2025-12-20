package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type UI struct {
	buttons map[string]*Button
}

type Button struct {
	text     string
	position rl.Vector2
	rX       int
	rY       int
	action   func()
	enabled  bool
}

func (g *Game) drawUI() {
	for _, v := range g.UI.buttons {
		if v.enabled {
			v.draw()
		}
	}
}

func NewButton(text string, positionX, positionY int, action func()) *Button {
	return &Button{
		text:     text,
		position: rl.Vector2{X: float32(positionX), Y: float32(positionY)},
		rX:       40,
		rY:       20,
		action:   action,
		enabled:  false,
	}
}

func (b *Button) MouseOnButton() bool {
	point := rl.GetMousePosition()
	// Avoid division by zero
	if b.rX <= 0 || b.rY <= 0 {
		return false
	}

	// Normalized squared distance from center
	dx := point.X - b.position.X
	dy := point.Y - b.position.Y

	termX := (dx * dx) / float32(b.rX*b.rX)
	termY := (dy * dy) / float32(b.rY*b.rY)

	// If sum is <= 1.0, point is inside or on edge
	return (termX + termY) <= 1.0
}

func (b *Button) draw() {
	rl.DrawEllipse(int32(b.position.X), int32(b.position.Y), float32(b.rX), float32(b.rY), rl.Black)
	textWidth := rl.MeasureText(b.text, 15)
	rl.DrawText(b.text, int32(b.position.X)-textWidth/2, int32(b.position.Y)-7, 15, rl.Blue)
}
