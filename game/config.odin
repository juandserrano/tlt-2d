package game

import "core:time"

GameConfig :: struct {
	lastModTime: time.Time,
	configPath:  string,
	GameName:    string,
	Window:      struct {
		Width:           int `json:"width"`,
		Height:          int `json:"height"`,
		BackgroundColor: struct {
			R: f32 `json:"r"`,
			G: f32 `json:"g"`,
			B: f32 `json:"b"`,
			A: f32 `json:"a"`,
		},
		TargetFPS:       int `json:"target_fps"`,
	},
	Camera:      struct {
		MoveSpeed: f32 `json:"move_speed"`,
	} `json:"camera"`,
	World:       struct {
		AnimateSun: bool `json:"animate_sun"`,
	} `json:"world"`,
	Rules:       struct {
		HandLimit:       int `json:"hand_limit"`,
		DeckComposition: struct {
			AttackPawnQty:   int `json:"attack_pawn_qty"`,
			AttackKnightQty: int `json:"attack_knight_qty"`,
			AttackBishopQty: int `json:"attack_bishop_qty"`,
			AttackQueenQty:  int `json:"attack_queen_qty"`,
			AttackKingQty:   int `json:"attack_king_qty"`,
		} `json:"deck_composition"`,
	} `json:"rules"`,
	Enemies:     struct {
		Pawn:   struct {
			Health: int `json:"health"`,
			Attack: int `json:"attack"`,
		} `json:"pawn"`,
		Knight: struct {
			Health: int `json:"health"`,
			Attack: int `json:"attack"`,
		} `json:"knight"`,
		Bishop: struct {
			Health: int `json:"health"`,
			Attack: int `json:"attack"`,
		} `json:"bishop"`,
		Queen:  struct {
			Health: int `json:"health"`,
			Attack: int `json:"attack"`,
		} `json:"queen"`,
		King:   struct {
			Health: int `json:"health"`,
			Attack: int `json:"attack"`,
		} `json:"king"`,
	} `json:"enemies"`,
	Player:      struct {
		Health: int `json:"health"`,
	} `json:"player"`,
	Animations:  struct {
		SlideSpeed: f32 `json:"slide_speed"`,
	} `json:"animations"`,
}
