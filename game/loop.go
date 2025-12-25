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

	g.toggleDebug()
	g.handleCamera()

	// Update Animations
	var activeAnims []*CardAnimation
	for _, anim := range g.cardAnimations {
		anim.Progress += dt * 2.5 // Adjust speed here
		if anim.Progress >= 1.0 {
			anim.OnFinish()
		} else {
			activeAnims = append(activeAnims, anim)
		}
	}
	g.cardAnimations = activeAnims

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
			if g.endingTurn {
				g.uiAlpha -= dt * 2.0
				if g.uiAlpha <= 0 {
					g.uiAlpha = 0
					g.endingTurn = false
					g.Turn = TurnComputer
					g.enemyMoveIndex = 0
					g.waitingForMoveAnimation = false
					g.waitingForSpawnAnimation = false
					for i := range g.playerHand.cards {
						g.playerHand.cards[i].selected = false
						g.playerHand.selectedCard = nil
					}
				}
			} else {
				g.uiAlpha += dt * 2.0
				if g.uiAlpha > 1.0 {
					g.uiAlpha = 1.0
				}
				g.TurnPlayer(dt)
			}

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
	var indices []int
	for i := range EnemiesInPlay {
		if EnemiesInPlay[i].currentHealth <= 0 {
			indices = append(indices, i)
		}
	}
	for i, idx := range indices {
		EnemiesInPlay[idx] = EnemiesInPlay[len(EnemiesInPlay)-(i+1)]
	}
	EnemiesInPlay = EnemiesInPlay[:len(EnemiesInPlay)-len(indices)]
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
	default:
	}
	if g.debugLevel != 0 {
		g.DrawWorldDebug()
	}
	rl.EndMode3D()

	// Draw Card Animations
	for _, anim := range g.cardAnimations {
		start := anim.StartPosition
		target3D := anim.TargetEnemy.visualPos
		targetScreen := rl.GetWorldToScreen(rl.Vector3{X: target3D.X, Y: 0.5, Z: target3D.Z}, g.camera) // Aim a bit higher

		pos := rl.Vector2Lerp(start, targetScreen, anim.Progress)
		scale := rl.Lerp(1.0, 0.2, anim.Progress)

		// Draw card at interpolated position, shrinking
		// We use a simplified draw logic here or reuse card.draw if careful with position
		// Let's manually draw texture for total control
		tex := anim.Card.texture

		// Optional: Rotate it as it flies
		rotation := anim.Progress * 360.0 * 2.0

		rl.DrawTextureEx(*tex, pos, rotation, scale, rl.White)
	}

	if g.Turn == TurnPlayer {
		g.drawCards()
		g.playerHand.draw(g.uiAlpha)
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
