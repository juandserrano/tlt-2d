package game

import "core:fmt"
import "core:time"
import rl "vendor:raylib"

GameState :: enum {
	StateMenu,
	StatePlaying,
	StatePause,
	StateGameOver,
}

TurnState :: enum {
	TurnPlayer,
	TurnResolving,
	TurnComputer,
}

Game :: struct {
	Config:                   GameConfig,
	State:                    GameState,
	Turn:                     TurnState,
	Round:                    Round,
	camera:                   rl.Camera3D,
	tiles:                    map[TileType]Tile,
	levels:                   map[int]Level,
	currentLevel:             int,
	debugLevel:               i8,
	enemyBag:                 EnemyBag,
	// playerHand   Hand
	// deck         Deck
	// discardPile    Deck
	playerCastle:             Castle,
	shaders:                  map[ShaderName]^rl.Shader,
	plainTileModel:           rl.Model,
	waterTileModel:           rl.Model,
	enemyModels:              map[EnemyType]^rl.Model,
	cardTextures:             map[CardType]^rl.Texture2D,
	sunLight:                 Light,
	spotLight:                Light,
	frameCount:               int,
	UI:                       UI,
	CameraShakeIntensity:     f32,
	enemyMoveIndex:           int,
	waitingForMoveAnimation:  bool,
	waitingForSpawnAnimation: bool,
	turnTransitionTimer:      f32,
	sounds:                   map[string]rl.Sound,
	music:                    map[string]rl.Music,
	// AnimationController      *AnimationController
	// ParticleManager          *ParticleManager
}

run :: proc() {
	game := Game{}
	init(&game)
	defer rl.CloseAudioDevice()
	defer rl.CloseWindow()
	Loop(&game)
}

init :: proc(g: ^Game) {
	// g.CheckAndLoadConfig(true)
	rl.SetConfigFlags({.WINDOW_RESIZABLE, .MSAA_4X_HINT})
	// rl.SetTargetFPS(g.Config.Window.TargetFPS)
	// rl.InitWindow(g.Config.Window.Width, g.Config.Window.Height, g.Config.GameName)
	rl.InitWindow(800, 600, "The Last Tower")
	rl.InitAudioDevice()
	g.debugLevel = 0
	initCamera(g)
	g.levels = make(map[int]Level)
	g.tiles = make(map[TileType]Tile)
	g.shaders = make(map[ShaderName]^rl.Shader)
	g.enemyModels = make(map[EnemyType]^rl.Model)
	g.cardTextures = make(map[CardType]^rl.Texture2D)
	g.UI.buttons = make(map[string]Button)
	g.sounds = make(map[string]rl.Sound)
	g.music = make(map[string]rl.Music)
	LoadResources(g)
	initShadersAndLights(g)
	// g.AnimationController = NewAnimationController()
	g.Round = NewRound(g)
	g.State = .StatePlaying
}

Loop :: proc(g: ^Game) {
	g.frameCount = 0
	for !rl.WindowShouldClose() {
		g.frameCount += 1
		if g.frameCount % 30 == 0 {
			// g.CheckAndLoadConfig(false)
		}
		Update(g)
		Draw(g)
	}

}

Update :: proc(g: ^Game) {
	dt := rl.GetFrameTime()
	rl.UpdateMusicStream(g.music["iron_at_the_gate"])
	toggleDebug(g)

	handleCamera(g)
	// g.ParticleManager.Update(dt)

	// // Update Animations
	// g.AnimationController.Update(dt)

	// // Update Enemies Animation via controller
	// g.AnimationController.UpdateEnemies(dt, EnemiesInPlay, g)

	// // TODO: Fin a better place for this
	// g.mouseOverEnemies()
	// /////

	#partial switch g.State {
	case .StatePlaying:
		if g.Round.TurnNumber == 0 {
			SetUpRound(&g.Round, g)
		}
		#partial switch g.Turn {
		case .TurnPlayer:
			// Fade in UI at start of player turn
			// if g.AnimationController.GetUIAlpha() < 1.0 {
			// Already fading in via AnimationController
			// }
			TurnPlayer(g, dt)

		// g.checkAndCleanEnemies()
		// case TurnResolving:
		// 	g.TurnResolve(dt)
		case .TurnComputer:
			TurnComputer(g, dt)

		case:
		}
	// case .StatePause:

	// case StateWorldEditor:
	case:

	}

	// g.UpdateShaders()
	// g.OnWindowSizeUpdate()

	if g.debugLevel > 0 {
		terminalDebug(g)
	}
}

Draw :: proc(g: ^Game) {
	rl.BeginDrawing()
	rl.ClearBackground(
		rl.Color {
			200,
			200,
			200,
			255,
			// u8(g.Config.Window.BackgroundColor.R * 255),
			// u8(g.Config.Window.BackgroundColor.G * 255),
			// u8(g.Config.Window.BackgroundColor.B * 255),
			// u8(g.Config.Window.BackgroundColor.A * 255),
		},
	)
	rl.BeginMode3D(GetRenderCamera(g))
	#partial switch g.State {
	case .StatePlaying:
		drawEnemies(g)
		fmt.printf("current level: %d", g.currentLevel)
		drawLevel(&g.levels[g.currentLevel])
		drawCastle(&g.playerCastle)
	// DrawParticles(&g.particleManager, g.camera)
	case:
		break
	}
	// if g.debugLevel != 0 {
	// 	g.DrawWorldDebug()
	// }
	rl.EndMode3D()

	// Draw Card Animations
	// g.AnimationController.DrawCardAttackAnimations(g.camera)

	// if g.Turn == TurnPlayer {
	// 	g.drawCards()
	// 	g.playerHand.draw(g.AnimationController.GetUIAlpha())
	// 	g.drawUI()
	// }
	// if g.debugLevel != 0 {
	// 	g.DrawStaticDebug()
	// }

	rl.EndDrawing()

}

nextTurn :: proc(g: ^Game) {
	// g.UpdateEnemies()
}
