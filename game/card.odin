package game

import "core:encoding/uuid"
import rl "vendor:raylib"

CardType :: enum {
	CardTypeAttackPawn,
	CardTypeAttackBisho,
	CardTypeAttackKnight,
	CardTypeAttackKing,
	CardTypeAttackQueen,
	CardTypeBack,
}

Deck :: struct {
	cards:    []^Card,
	position: rl.Vector2,
	canDraw:  bool,
}
Card :: struct {
	id:               uuid.Identifier,
	name:             string,
	texture:          ^rl.Texture2D,
	position:         rl.Vector2,
	available:        bool,
	selected:         bool,
	selectedOffset:   rl.Vector2,
	selectedRotation: f32,
	isShowing:        bool,
	backTexture:      ^rl.Texture2D,
	cardType:         CardType,
	positionInHand:   int,
}
