package game

import rl "github.com/gen2brain/raylib-go/raylib"

type UI struct {
	buttons []Button
}

type Button struct {
	text     string
	position rl.Vector2
	rX       int
	rY       int
}

func (g *Game) drawUI() {
	for i := range g.UI.buttons {
		g.UI.buttons[i].draw()
	}
}

func NewButton(text string, positionX, positionY int) Button {
	return Button{
		text:     text,
		position: rl.Vector2{X: float32(positionX), Y: float32(positionY)},
		rX:       40,
		rY:       20,
	}
}

func (b *Button) draw() {
	rl.DrawEllipseLines(int32(b.position.X), int32(b.position.Y), float32(b.rX), float32(b.rY), rl.Black)
	rl.DrawText(b.text, int32(b.position.X), int32(b.position.Y), 15, rl.Blue)
}
