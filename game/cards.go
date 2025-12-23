package game

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

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
	cards    []*Card
	position rl.Vector2
	canDraw  bool
}
type Card struct {
	id               uuid.UUID
	name             string
	texture          *rl.Texture2D
	position         rl.Vector2
	available        bool
	selected         bool
	selectedOffset   rl.Vector2
	selectedRotation float32
	isShowing        bool
	backTexture      *rl.Texture2D
	cardType         CardType
	positionInHand   int
}

func (g *Game) NewCard(cardType CardType, pos rl.Vector2, available bool) Card {
	c := Card{
		id:               uuid.New(),
		texture:          g.cardTextures[cardType],
		position:         pos,
		available:        available,
		selected:         false,
		selectedOffset:   rl.Vector2{X: 0, Y: -30},
		selectedRotation: -3.0,
		isShowing:        true,
		backTexture:      g.cardTextures[CardTypeBack],
		cardType:         cardType,
		name:             cardType.String(),
	}
	return c
}

func (c *Card) move(newPos rl.Vector2) {
	c.position = newPos
}

func (g *Game) NewDeck() Deck {
	var d Deck
	d.canDraw = true
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
		c.isShowing = false
		d.cards = append(d.cards, &c)
	}

	g.ShuffleCards(d.cards)
	for i := range d.cards {
		offset := float32(i) * 0.3
		d.cards[i].position.Y = d.position.Y - offset
		d.cards[i].position.X = d.position.X + offset
	}
	return d
}

func (g *Game) ShuffleCards(slice []*Card) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func (g *Game) drawCards() {
	for i := range g.deck.cards {
		g.deck.cards[i].draw(1)
	}
}

func (c *Card) update() {
	if c.selected {

	}
}

func (g *Game) updateCards() {

}

func (c *Card) draw(scale float32) {
	offset := rl.Vector2{}
	var rotation float32 = 0.0
	if c.selected {
		offset = c.selectedOffset
		rotation = c.selectedRotation
	}
	if c.isShowing {
		rl.DrawTextureEx(*c.texture, rl.Vector2Add(c.position, offset), rotation, scale, rl.White)

	} else {
		rl.DrawTextureEx(*c.backTexture, rl.Vector2Add(c.position, offset), 0, scale, rl.White)
	}
}

func (c *Card) toggleSelected() {
	c.selected = !c.selected
}

func (g *Game) drawToTopHand(h *Hand) {
	if g.deck.canDraw {
		availablePositions := 0
		for i := range h.cardPositions {
			if h.cardPositions[i].available {
				availablePositions++
			}
		}
		for range availablePositions {
			err := g.deck.moveTopCardToHand(h)
			if err != nil {
				return
			}
		}
		g.deck.canDraw = false

	}

}

func (d *Deck) moveTopCardToHand(h *Hand) error {
	if len(d.cards) == 0 {
		return fmt.Errorf("no more cards on deck")
	}
	pos, worldPos, err := h.nextAvailablePosition()
	if err != nil {
		return err
	}
	h.cards = append(h.cards, d.cards[len(d.cards)-1]) // Add card to hand cards
	// Move position of card to deck position
	newCardInHand := h.cards[len(h.cards)-1]
	h.cardPositions[pos].available = false
	newCardInHand.isShowing = true
	newCardInHand.positionInHand = pos
	newCardInHand.position = rl.Vector2{
		X: worldPos.X - float32(newCardInHand.texture.Width)/2.0*(float32(pos)+1),
		Y: worldPos.Y - float32(newCardInHand.texture.Height)/2.0}
	d.cards = d.cards[:len(d.cards)-1] // Remove card from deck

	return nil
}

func (c *Card) isMouseOnCard() bool {
	mousePos := rl.GetMousePosition()
	var bounds rl.Rectangle
	if c.selected {
		bounds = rl.NewRectangle(c.position.X+c.selectedOffset.X, c.position.Y+c.selectedOffset.Y, float32(c.texture.Width), float32(c.texture.Height))

	} else {
		bounds = rl.NewRectangle(c.position.X, c.position.Y, float32(c.texture.Width), float32(c.texture.Height))

	}
	return rl.CheckCollisionPointRec(mousePos, bounds)
}

func (c *Card) addToplay(g *Game) {
	g.cardsToPlay = append(g.cardsToPlay, c)
}

func (c *Card) attackEnemy(enemy *Enemy, h *Hand) {
	h.moveCardToDiscardPile(c)
	if enemy.enemyType == EnemyTypeKnight {
		if rand.Float32() < 0.2 {
			enemy.move()
			return
		}
	}
	enemy.currentHealth--
}

func (t CardType) String() string {
	switch t {
	case CardTypeAttackBishop:
		return "Attack Bishop"
	case CardTypeAttackKnight:
		return "Attack Knight"
	case CardTypeAttackPawn:
		return "Attack Pawn"
	case CardTypeAttackQueen:
		return "Attack Queen"
	case CardTypeAttackKing:
		return "Attack King"
	default:
		return ""
	}
}
