package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	texture           rl.Texture2D
	position          rl.Vector2
	prevPos           rl.Vector2
	targetPos         rl.Vector2
	gridX             int
	gridY             int
	xOffset           float32
	yOffset           float32
	isMoving          bool
	moveSpeed         float32
	moveCooldown      float32
	moveCooldownTimer float32
	moveOnCooldown    bool
}

func (p *Player) draw() {
	rl.DrawTextureEx(p.texture, p.position, 0, 0.15, rl.White)
}

func (g *Game) initPlayer() {
	g.player = Player{texture: rl.LoadTexture("assets/pawn2.png"), position: GetTileCenter(0, 0, g.tex)}
	g.player.xOffset = -float32(g.player.texture.Width) * 0.15 / 2
	g.player.yOffset = -float32(g.player.texture.Height) * 0.15
	g.player.position.X += g.player.xOffset
	g.player.position.Y += g.player.yOffset
	g.player.moveCooldown = 0.3
	g.player.moveCooldownTimer = g.player.moveCooldown
	g.player.moveOnCooldown = false
	g.player.moveSpeed = 20
}

func (g *Game) MovePlayer(dx, dy int) {
	if g.player.isMoving {
		return
	}

	g.player.gridX += dx
	g.player.gridY += dy
	if (g.player.gridX) < 0 {
		g.player.gridX = 0
	}
	if (g.player.gridY) < 0 {
		g.player.gridY = 0
	}

	l := g.levels[g.currentLevel]
	newPos := l.GetTileCenterPosition(g.player.gridX, g.player.gridY)
	g.player.prevPos = g.player.position
	g.player.targetPos.X = newPos.X + g.player.xOffset
	g.player.targetPos.Y = newPos.Y + g.player.yOffset

	g.player.isMoving = true

}

func (p *Player) update(dt float32) {
	if !p.isMoving {
		return
	}

	// Math: Move RenderPos towards TargetPos
	// The '10.0' here is the speed. Higher = Faster.
	// 'dt' is DeltaTime (rl.GetFrameTime()) to ensure smooth movement at any FPS.

	p.position = rl.Vector2Lerp(
		p.position,
		p.targetPos,
		p.moveSpeed*dt,
	)

	// Snap to position when very close (to stop micro-jittering)
	dist := rl.Vector2Distance(p.position, p.targetPos)
	if dist < 5.0 {
		p.position = p.targetPos
		p.isMoving = false
	}

}
