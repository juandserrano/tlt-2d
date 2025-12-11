package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) handleInput(dt float32) {
	g.handleCamera()
	g.handlePlayerMovement(dt)
	g.toggleDebug()
}

func (g *Game) handlePlayerMovement(dt float32) {

	if g.player.moveOnCooldown {
		g.player.moveCooldownTimer -= dt
		if g.player.moveCooldownTimer <= 0 {
			g.player.moveOnCooldown = false
			g.player.moveCooldownTimer = g.player.moveCooldown
		}
	}

	if rl.IsKeyDown(rl.KeyW) && !g.player.moveOnCooldown {
		g.MovePlayer(0, -1)
		g.player.moveOnCooldown = true
	}
	if rl.IsKeyDown(rl.KeyUp) && !g.player.moveOnCooldown {
		g.MovePlayer(0, -1)
		g.player.moveOnCooldown = true
	}
	if rl.IsKeyDown(rl.KeyS) && !g.player.moveOnCooldown {
		g.MovePlayer(0, 1)
		g.player.moveOnCooldown = true
	}
	if rl.IsKeyDown(rl.KeyDown) && !g.player.moveOnCooldown {
		g.MovePlayer(0, 1)
		g.player.moveOnCooldown = true
	}
	if rl.IsKeyDown(rl.KeyA) && !g.player.moveOnCooldown {
		g.MovePlayer(-1, 0)
		g.player.moveOnCooldown = true
	}
	if rl.IsKeyDown(rl.KeyLeft) && !g.player.moveOnCooldown {
		g.MovePlayer(-1, 0)
		g.player.moveOnCooldown = true
	}
	if rl.IsKeyDown(rl.KeyD) && !g.player.moveOnCooldown {
		g.MovePlayer(1, 0)
		g.player.moveOnCooldown = true
	}
	if rl.IsKeyDown(rl.KeyRight) && !g.player.moveOnCooldown {
		g.MovePlayer(1, 0)
		g.player.moveOnCooldown = true
	}
}
