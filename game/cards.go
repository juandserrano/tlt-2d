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

func (g *Game) NewDiscardPile() Deck {
	var dp Deck
	dp.position = rl.Vector2{X: float32(g.Config.Window.Width) - float32(g.cardTextures[CardTypeAttackPawn].Width) - 10, Y: 20}
	return dp
}

func (d *Deck) Pop() (*Card, error) {
	if len(d.cards) == 0 {
		return nil, fmt.Errorf("deck is empty")
	}
	lastIdx := len(d.cards) - 1
	card := d.cards[lastIdx]
	d.cards = d.cards[:lastIdx]
	return card, nil
}

func (d *Deck) Push(c *Card) {
	d.cards = append(d.cards, c)
}

func (g *Game) NewDeck() Deck {
	var d Deck
	d.canDraw = true
	d.position = rl.Vector2{X: 10, Y: 20}

	counts := map[CardType]int{
		CardTypeAttackPawn:   g.Config.Rules.DeckComposition.AttackPawnQty,
		CardTypeAttackKnight: g.Config.Rules.DeckComposition.AttackKnightQty,
		CardTypeAttackBishop: g.Config.Rules.DeckComposition.AttackBishopQty,
		CardTypeAttackQueen:  g.Config.Rules.DeckComposition.AttackQueenQty,
		CardTypeAttackKing:   g.Config.Rules.DeckComposition.AttackKingQty,
	}

	for cardType, count := range counts {
		for range count {
			c := g.NewCard(cardType, d.position, true)
			c.isShowing = false
			d.cards = append(d.cards, &c)
		}
	}

	g.ShuffleCards(d.cards)

	// Initial positioning
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
		g.deck.cards[i].draw(1, g.AnimationController.GetUIAlpha())
	}
	for i := range g.discardPile.cards {
		g.discardPile.cards[i].draw(1, g.AnimationController.GetUIAlpha())
	}
}

func (c *Card) update() {
	if c.selected {

	}
}

func (g *Game) updateCards() {

}

func (c *Card) draw(scale float32, alpha float32) {
	offset := rl.Vector2{}
	var rotation float32 = 0.0
	if c.selected {
		offset = c.selectedOffset
		rotation = c.selectedRotation
	}
	color := rl.White
	color.A = uint8(alpha * 255)
	if c.isShowing {
		if c.selected {
			rl.BeginShaderMode(outlineShader)
			rl.DrawTextureEx(*c.texture, rl.Vector2Add(c.position, offset), rotation, scale, color)
			rl.EndShaderMode()

		} else {
			rl.DrawTextureEx(*c.texture, rl.Vector2Add(c.position, offset), rotation, scale, color)
		}
	} else {
		rl.DrawTextureEx(*c.backTexture, rl.Vector2Add(c.position, offset), 0, scale, color)
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
	card, err := d.Pop()
	if err != nil {
		return err
	}

	pos, worldPos, err := h.nextAvailablePosition()
	if err != nil {
		// If hand is full, we need to return the card to the deck (conceptually).
		// For now, since Pop removed it, we push it back.
		d.Push(card)
		return err
	}

	h.cards = append(h.cards, card) // Add card to hand cards

	// Move position of card to deck position
	h.cardPositions[pos].available = false
	card.isShowing = true
	card.positionInHand = pos
	card.position = rl.Vector2{
		X: worldPos.X - float32(card.texture.Width)/2.0*(float32(pos)+1),
		Y: worldPos.Y - float32(card.texture.Height)/2.0}

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

// func (c *Card) addToplay(g *Game) {
// 	g.cardsToPlay = append(g.cardsToPlay, c)
// }

func (g *Game) StartCardAttack(card *Card, enemy *Enemy) {
	g.playerHand.moveCardToDiscardPile(card, &g.discardPile, g)

	onFinish := func() {
		if enemy.enemyType == EnemyTypeKnight {
			if rand.Float32() < 0.2 {
				enemy.move()
				return
			}
		}
		g.ParticleManager.SpawnExplosion(enemy.visualPos, 10, rl.Blue)
		g.ParticleManager.SpawnExplosion(enemy.visualPos, 20, rl.Yellow)
		g.ParticleManager.SpawnExplosion(enemy.visualPos, 5, rl.Red)
		enemy.currentHealth--
	}

	anim := &CardAnimation{
		Card:          card,
		StartPosition: card.position,
		TargetEnemy:   enemy,
		Progress:      0.0,
		OnFinish:      onFinish,
	}
	g.AnimationController.AddCardAttackAnimation(anim)
}

func CanAttack(card *Card, enemy *Enemy) bool {
	switch card.cardType {
	case CardTypeAttackPawn:
		return enemy.enemyType == EnemyTypePawn
	case CardTypeAttackBishop:
		return enemy.enemyType == EnemyTypeBishop
	case CardTypeAttackKnight:
		return enemy.enemyType == EnemyTypeKnight
	case CardTypeAttackQueen:
		return enemy.enemyType == EnemyTypeQueen
	case CardTypeAttackKing:
		return enemy.enemyType == EnemyTypeKing
	default:
		return false
	}
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
