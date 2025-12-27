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

func (g *Game) mouseOverEnemies() {
	if g.frameCount%5 == 0 {
		ray := rl.GetScreenToWorldRay(rl.GetMousePosition(), g.camera)
		for i := range EnemiesInPlay {
			bb := rl.GetModelBoundingBox(*EnemiesInPlay[i].model)
			pos := g.GetTileCenter(EnemiesInPlay[i].gridPos)
			bb.Min = rl.Vector3Add(bb.Min, pos)
			bb.Max = rl.Vector3Add(bb.Max, pos)

			rayCollision := rl.GetRayCollisionBox(ray, bb)
			if rayCollision.Hit {
				EnemiesInPlay[i].healthBarShowing = true
			} else {
				EnemiesInPlay[i].healthBarShowing = false
			}
		}
	}
}

func (g *Game) Update(dt float32) {
	rl.UpdateMusicStream(g.music["iron_at_the_gate"])
	g.toggleDebug()

	g.handleCamera()
	g.ParticleManager.Update(dt)

	// Update Animations
	g.AnimationController.Update(dt)

	// Update Enemies Animation
	for i := range EnemiesInPlay {
		EnemiesInPlay[i].Update(dt, g)
	}

	// TODO: Fin a better place for this
	g.mouseOverEnemies()
	/////

	switch g.State {
	case StatePlaying:
		if g.Round.TurnNumber == 0 {
			g.Round.SetUp(g)
		}
		switch g.Turn {
		case TurnPlayer:
			// Fade in UI at start of player turn
			if g.AnimationController.GetUIAlpha() < 1.0 {
				// Already fading in via AnimationController
			}
			g.TurnPlayer(dt)

			g.checkAndCleanEnemies()
		// case TurnResolving:
		// 	g.TurnResolve(dt)
		case TurnComputer:
			g.TurnComputer(dt)
		}
	case StatePause:

		// case StateWorldEditor:

	}

	g.UpdateShaders()
	g.OnWindowSizeUpdate()

	if g.debugLevel > 0 {
		g.TerminalDebug()

	}
}

func (g *Game) checkAndCleanEnemies() {
	n := 0
	for _, e := range EnemiesInPlay {
		if e.currentHealth > 0 {
			EnemiesInPlay[n] = e
			n++
		}
	}
	EnemiesInPlay = EnemiesInPlay[:n]
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(color.RGBA{uint8(g.Config.Window.BackgroundColor.R * 255), uint8(g.Config.Window.BackgroundColor.G * 255), uint8(g.Config.Window.BackgroundColor.B * 255), uint8(g.Config.Window.BackgroundColor.A * 255)})
	rl.BeginMode3D(g.GetRenderCamera())
	switch g.State {
	// case StateWorldEditor:
	// 	g.DrawLevel(g.currentLevel)
	case StatePlaying:
		g.drawEnemies()
		g.DrawLevel(g.currentLevel)
		g.playerCastle.draw()
		g.ParticleManager.Draw(g.camera)
	default:
	}
	if g.debugLevel != 0 {
		g.DrawWorldDebug()
	}
	rl.EndMode3D()

	// Draw Card Animations
	g.AnimationController.DrawCardAttackAnimations(g.camera)

	if g.Turn == TurnPlayer {
		g.drawCards()
		g.playerHand.draw(g.AnimationController.GetUIAlpha())
		g.drawUI()
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
