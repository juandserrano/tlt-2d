package game

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Loop() {
	g.frameCount = 0
	for !rl.WindowShouldClose() {
		g.frameCount++
		if g.frameCount%30 == 0 {
			g.CheckAndLoadConfig(false)
		}
		dt := rl.GetFrameTime()
		g.Update(dt)
		g.Draw()
	}
}

func (g *Game) Update(dt float32) {

	switch g.State {
	case StatePlaying:
		switch g.Turn {
		case TurnPlayer:
			g.TurnPlayer(dt)
		case TurnResolving:
			g.TurnResolve(dt)
		case TurnComputer:
			g.TurnComputer(dt)
		}
	case StatePause:

	}

	g.UpdateShaders()
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(color.RGBA{uint8(g.Config.Window.BackgroundColor.R * 255), uint8(g.Config.Window.BackgroundColor.G * 255), uint8(g.Config.Window.BackgroundColor.B * 255), uint8(g.Config.Window.BackgroundColor.A * 255)})
	rl.BeginMode3D(g.camera)
	rl.DrawGrid(10, 1.0)
	g.DrawLevel(g.currentLevel)
	g.playerCastle.draw()
	g.drawEnemies()
	if g.debug {
		g.DrawWorldDebug()
	}
	rl.EndMode3D()
	if g.debug {
		g.DrawStaticDebug()
	}
	rl.EndDrawing()
}

func (g *Game) DrawLevel(currentLevel int) {
	level := g.levels[currentLevel]
	level.Draw()
}
