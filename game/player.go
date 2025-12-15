package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Castle struct {
	model    rl.Model
	position rl.Vector3
	gridX    int
	gridZ    int
}

func (p *Castle) draw() {
	rl.DrawModel(p.model, p.position, 1.0, rl.White)
}

func (g *Game) initPlayerCastle() {
	g.playerCastle.gridX = 0
	g.playerCastle.gridZ = 0
	startXZPos := GridToWorldHex(g.playerCastle.gridX, g.playerCastle.gridZ, HEX_TILE_WIDTH/2.0)
	g.playerCastle.position.X = startXZPos.X
	g.playerCastle.position.Z = startXZPos.Y
	g.playerCastle.position.Y = 0
}

func (p *Castle) update(dt float32) {

}
