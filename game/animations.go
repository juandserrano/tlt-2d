package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

// CardAnimation represents a card flying towards and attacking an enemy
type CardAnimation struct {
	Card          *Card
	StartPosition rl.Vector2
	TargetEnemy   *Enemy
	Progress      float32
	OnFinish      func()
}

// CardSlideAnimation represents a card sliding to a new position in hand
type CardSlideAnimation struct {
	Card           *Card
	StartPosition  rl.Vector2
	TargetPosition rl.Vector2
	Progress       float32
}

// AnimationController manages all animations in the game
type AnimationController struct {
	cardAttackAnimations []*CardAnimation
	cardSlideAnimations  []*CardSlideAnimation
}

// NewAnimationController creates a new animation controller
func NewAnimationController() *AnimationController {
	return &AnimationController{
		cardAttackAnimations: []*CardAnimation{},
		cardSlideAnimations:  []*CardSlideAnimation{},
	}
}

// AddCardAttackAnimation adds a card attack animation to the controller
func (ac *AnimationController) AddCardAttackAnimation(anim *CardAnimation) {
	ac.cardAttackAnimations = append(ac.cardAttackAnimations, anim)
}

// AddCardSlideAnimation adds a card slide animation to the controller
func (ac *AnimationController) AddCardSlideAnimation(anim *CardSlideAnimation) {
	ac.cardSlideAnimations = append(ac.cardSlideAnimations, anim)
}

// Update processes all active animations
func (ac *AnimationController) Update(dt float32) {
	// Update card attack animations
	var activeAttackAnims []*CardAnimation
	for _, anim := range ac.cardAttackAnimations {
		anim.Progress += dt * 2.5 // Adjust speed here
		if anim.Progress >= 1.0 {
			anim.OnFinish()
		} else {
			activeAttackAnims = append(activeAttackAnims, anim)
		}
	}
	ac.cardAttackAnimations = activeAttackAnims

	// Update card slide animations
	var activeSlideAnims []*CardSlideAnimation
	for _, anim := range ac.cardSlideAnimations {
		anim.Progress += dt * 5.0 // Fast, responsive sliding
		if anim.Progress >= 1.0 {
			anim.Card.position = anim.TargetPosition
		} else {
			anim.Card.position = rl.Vector2Lerp(anim.StartPosition, anim.TargetPosition, anim.Progress)
			activeSlideAnims = append(activeSlideAnims, anim)
		}
	}
	ac.cardSlideAnimations = activeSlideAnims
}

// DrawCardAttackAnimations renders all active card attack animations
func (ac *AnimationController) DrawCardAttackAnimations(camera rl.Camera3D) {
	for _, anim := range ac.cardAttackAnimations {
		start := anim.StartPosition
		target3D := anim.TargetEnemy.visualPos
		targetScreen := rl.GetWorldToScreen(rl.Vector3{X: target3D.X, Y: 0.5, Z: target3D.Z}, camera)

		pos := rl.Vector2Lerp(start, targetScreen, anim.Progress)
		scale := rl.Lerp(1.0, 0.2, anim.Progress)

		// Draw card at interpolated position, shrinking
		tex := anim.Card.texture

		// Rotate it as it flies
		rotation := anim.Progress * 360.0 * 2.0

		rl.DrawTextureEx(*tex, pos, rotation, scale, rl.White)
	}
}

// IsCardAnimating checks if a specific card is currently in a slide animation
func (ac *AnimationController) IsCardAnimating(cardID uuid.UUID) bool {
	for _, anim := range ac.cardSlideAnimations {
		if anim.Card.id == cardID {
			return true
		}
	}
	return false
}
