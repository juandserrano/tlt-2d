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

	g.toggleDebug()
	g.handleCamera()
	switch g.State {
	case StatePlaying:
		if g.Round.TurnNumber == 0 {
			g.Round.SetUp(g)
		}
		switch g.Turn {
		case TurnPlayer:
			g.TurnPlayer(dt)
		case TurnResolving:
			g.TurnResolve(dt)
		case TurnComputer:
			g.TurnComputer(dt)
		}
	case StatePause:

	case StateWorldEditor:

	}

	g.UpdateShaders()
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(color.RGBA{uint8(g.Config.Window.BackgroundColor.R * 255), uint8(g.Config.Window.BackgroundColor.G * 255), uint8(g.Config.Window.BackgroundColor.B * 255), uint8(g.Config.Window.BackgroundColor.A * 255)})
	rl.BeginMode3D(g.camera)
	switch g.State {
	case StateWorldEditor:
		g.DrawLevel(g.currentLevel)
	case StatePlaying:
		g.drawEnemies()
		g.DrawLevel(g.currentLevel)
		g.playerCastle.draw()
	default:
	}
	if g.debugLevel != 0 {
		g.DrawWorldDebug()
	}
	rl.EndMode3D()
	if g.Turn == TurnPlayer {
		g.drawCards()
		g.playerHand.draw()

	}
	if g.debugLevel != 0 {
		g.DrawStaticDebug()
	}

	rl.EndDrawing()
}

func (g *Game) DrawLevel(currentLevel int) {
	level := g.levels[currentLevel]
	level.Draw()
}
