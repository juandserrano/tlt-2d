package game

import rl "vendor:raylib"

Round :: struct {
	TurnNumber: int,
}

NewRound :: proc(g: ^Game) -> Round {
	return Round{TurnNumber = 0}
}

SetUpRound :: proc(r: ^Round, g: ^Game) {
	// g.enemyBag = g.NewEnemyBag()
	// g.playerHand = g.NewHand()
	// g.discardPile = g.NewDiscardPile()
	// g.deck = g.NewDeck()
	// g.UI.buttons["draw"] = NewButton("draw", 300, 100, proc() {drawToTopHand(g, &g.playerHand)})
	// g.UI.buttons["end_turn"] = NewButton("End Turn", 300, 300, proc() {actionEndTurn(g)})
	g.Turn = .TurnPlayer
	// g.LoadLevelTiles(1)
	initPlayerCastle(g)
	// startingEnemies := g.enemyBag.PickStartingEnemies()

	// g.spawnSetUpEnemies(startingEnemies)

	// if rl.IsMusicStreamPlaying(g.music["iron_at_the_gate"]) {
	// 	rl.StopMusicStream(g.music["iron_at_the_gate"])
	// }
	// rl.PlayMusicStream(g.music["iron_at_the_gate"])
	g.Round.TurnNumber = 1

}
