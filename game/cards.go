package game

import rl "github.com/gen2brain/raylib-go/raylib"

type CardType int

const (
	CardTypeAttackPawn CardType = iota
	CardTypeAttackBishop
	CardTypeAttackKnight
	CardTypeAttackKing
	CardTypeAttackQueen
)

type Deck struct {
	cards []Card
}
type Card struct {
	model     *rl.Model
	material  *rl.Material
	position  rl.Vector3
	available bool
}

func (g *Game) NewCard(cardType CardType, pos rl.Vector3, available bool) Card {
	c := Card{
		model:     g.cardModels[cardType],
		position:  pos,
		available: available,
	}
	return c
}

func (c *Card) move(newPos rl.Vector3) {
	c.position = newPos
}

func (g *Game) NewDeck() Deck {
	var d Deck
	for i := range 30 {
		c := g.NewCard(CardTypeAttackPawn, rl.Vector3{X: 2, Y: 2 + (float32(i) * 0.012), Z: 3}, true)
		d.cards = append(d.cards, c)
	}
	return d
}

func (g *Game) drawCards() {
	for i := range g.deck.cards {
		g.deck.cards[i].draw()
	}
}

func (c *Card) draw() {
	rl.DrawModelEx(*c.model, c.position,
		rl.Vector3{0, 1, 0}, 0, rl.Vector3{1, 1, 1}, rl.White)
}
