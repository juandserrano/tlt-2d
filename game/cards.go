package game

import rl "github.com/gen2brain/raylib-go/raylib"

type CardType int

const (
	CardTypeAttackPawn CardType = iota
	CardTypeAttackBishop
	CardTypeAttackKnight
	CardTypeAttackKing
	CardTypeAttackQueen
	CardTypeBack
)

type Deck struct {
	cards    []Card
	position rl.Vector2
}
type Card struct {
	texture         *rl.Texture2D
	position        rl.Vector2
	available       bool
	selected        bool
	selectedYOffset int
	isShowing       bool
	backTexture     *rl.Texture2D
}

func (g *Game) NewCard(cardType CardType, pos rl.Vector2, available bool) Card {
	c := Card{
		texture:         g.cardTextures[cardType],
		position:        pos,
		available:       available,
		selected:        false,
		selectedYOffset: 10,
		isShowing:       false,
		backTexture:     g.cardTextures[CardTypeBack],
	}
	return c
}

func (c *Card) move(newPos rl.Vector2) {
	c.position = newPos
}

func (g *Game) NewDeck() Deck {
	var d Deck
	d.position = rl.Vector2{X: 10, Y: 20}
	totalCards := g.Config.Rules.DeckComposition.AttackPawnQty + g.Config.Rules.DeckComposition.AttackKnightQty + g.Config.Rules.DeckComposition.AttackBishopQty
	pawnLeft := g.Config.Rules.DeckComposition.AttackPawnQty
	knightLeft := g.Config.Rules.DeckComposition.AttackKnightQty
	bishopLeft := g.Config.Rules.DeckComposition.AttackBishopQty
	for i := range totalCards {
		offset := float32(i) * 0.3
		var c Card
		if pawnLeft > 0 {
			c = g.NewCard(CardTypeAttackPawn, rl.Vector2{X: d.position.X + offset, Y: d.position.Y - offset}, true)
			pawnLeft--
		} else if knightLeft > 0 {
			c = g.NewCard(CardTypeAttackKnight, rl.Vector2{X: d.position.X + offset, Y: d.position.Y - offset}, true)
			knightLeft--
		} else if bishopLeft > 0 {
			c = g.NewCard(CardTypeAttackBishop, rl.Vector2{X: d.position.X + offset, Y: d.position.Y - offset}, true)
			bishopLeft--
		}
		d.cards = append(d.cards, c)
	}
	return d
}

func (g *Game) drawCards() {
	for i := range g.deck.cards {
		g.deck.cards[i].draw()
	}
}

func (c *Card) update() {
	if c.selected {

	}
}

func (g *Game) updateCards() {

}

func (c *Card) draw() {
	offset := 0
	if c.selected {
		offset = c.selectedYOffset
	}
	if c.isShowing {
		rl.DrawTexture(*c.texture, int32(c.position.X), int32(c.position.Y-float32(offset)), rl.White)

	} else {
		rl.DrawTexture(*c.backTexture, int32(c.position.X), int32(c.position.Y-float32(offset)), rl.White)
	}
}

func (d *Deck) toggleSelectTopCard() {
	d.cards[len(d.cards)-1].selected = !d.cards[len(d.cards)-1].selected
}

func (d *Deck) moveTopCardToHand(h *Hand) {
	availablePosInHand, err := h.nextAvailablePosition()
	if err != nil {
		return
	}
	h.cards = append(h.cards, d.cards[len(d.cards)-1]) // Add card to hand cards
	// Move position of card to deck position
	newCardInHand := &h.cards[len(h.cards)-1]
	newCardInHand.position = rl.Vector2{
		X: availablePosInHand.X - float32(newCardInHand.texture.Width)/2.0,
		Y: availablePosInHand.Y - float32(newCardInHand.texture.Height)/2.0}
	d.cards = d.cards[:len(d.cards)-1] // Remove card from deck

}

func (d *Deck) isMouseOnTopCard() bool {
	mousePos := rl.GetMousePosition()
	topCard := d.cards[len(d.cards)-1]
	var bounds rl.Rectangle
	if topCard.selected {
		bounds = rl.NewRectangle(topCard.position.X, topCard.position.Y-float32(topCard.selectedYOffset), float32(topCard.texture.Width), float32(topCard.texture.Height))

	} else {
		bounds = rl.NewRectangle(topCard.position.X, topCard.position.Y, float32(topCard.texture.Width), float32(topCard.texture.Height))

	}
	return rl.CheckCollisionPointRec(mousePos, bounds)
}
