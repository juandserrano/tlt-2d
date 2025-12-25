package game

import (
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func TestDeckPop(t *testing.T) {
	d := Deck{
		cards: []*Card{
			{name: "Card1"},
			{name: "Card2"},
		},
	}

	// First pop
	c, err := d.Pop()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if c.name != "Card2" {
		t.Errorf("Expected Card2, got %s", c.name)
	}
	if len(d.cards) != 1 {
		t.Errorf("Expected 1 card remaining, got %d", len(d.cards))
	}

	// Second pop
	c, err = d.Pop()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if c.name != "Card1" {
		t.Errorf("Expected Card1, got %s", c.name)
	}
	if len(d.cards) != 0 {
		t.Errorf("Expected 0 cards remaining, got %d", len(d.cards))
	}

	// Third pop (empty)
	_, err = d.Pop()
	if err == nil {
		t.Error("Expected error on empty deck, got nil")
	}
}

func TestDeckPush(t *testing.T) {
	d := Deck{}
	c := &Card{name: "TestCard"}

	d.Push(c)

	if len(d.cards) != 1 {
		t.Errorf("Expected 1 card, got %d", len(d.cards))
	}
	if d.cards[0] != c {
		t.Error("Card pointer mismatch")
	}
}

func TestNewDeck(t *testing.T) {
	// Mock config
	var cfg GameConfig
	// Initialize nested anonymous structs using zero-value assignment trick
	cfg.Rules.DeckComposition.AttackPawnQty = 2
	cfg.Rules.DeckComposition.AttackKnightQty = 1
	cfg.Rules.DeckComposition.AttackBishopQty = 0
	cfg.Rules.DeckComposition.AttackQueenQty = 0
	cfg.Rules.DeckComposition.AttackKingQty = 0

	g := &Game{
		Config:       cfg,
		cardTextures: make(map[CardType]*rl.Texture2D),
	}

	// Mock textures to avoid panic in NewCard (nil map access)
	// NewCard uses g.cardTextures[cardType]

	d := g.NewDeck()

	expectedTotal := 3
	if len(d.cards) != expectedTotal {
		t.Errorf("Expected %d cards, got %d", expectedTotal, len(d.cards))
	}
}
