package game

import "core:encoding/uuid"
import "core:math/linalg"
import rl "vendor:raylib"

// CardAnimation represents a card flying towards and attacking an enemy
CardAnimation :: struct {
	Card:          ^Card,
	StartPosition: rl.Vector2,
	TargetEnemy:   ^Enemy,
	Progress:      f32,
	OnFinish:      proc(),
}

// CardSlideAnimation represents a card sliding to a new position in hand
CardSlideAnimation :: struct {
	Card:           ^Card,
	StartPosition:  rl.Vector2,
	TargetPosition: rl.Vector2,
	Progress:       f32,
}

// UIFadeAnimation represents UI fading in or out
UIFadeAnimation :: struct {
	TargetAlpha: f32, // 0.0 for fade out, 1.0 for fade in
	Speed:       f32, // Speed multiplier
	OnComplete:  proc(), // Optional callback when fade completes
}

// AnimationController manages all animations in the game
AnimationController :: struct {
	cardAttackAnimations: [dynamic]^CardAnimation,
	cardSlideAnimations:  [dynamic]^CardSlideAnimation,
	uiFadeAnimation:      ^UIFadeAnimation,
	uiAlpha:              f32, // Current UI alpha value
}

// NewAnimationController creates a new animation controller
NewAnimationController :: proc() -> AnimationController {
	return AnimationController {
		cardAttackAnimations = [dynamic]^CardAnimation{},
		cardSlideAnimations  = [dynamic]^CardSlideAnimation{},
		uiAlpha              = 1.0, // Start with UI fully visible
	}
}

// AddCardAttackAnimation adds a card attack animation to the controller
AddCardAttackAnimation :: proc(ac: ^AnimationController, anim: ^CardAnimation) {
	append(&ac.cardAttackAnimations, anim)
}

// AddCardSlideAnimation adds a card slide animation to the controller
AddCardSlideAnimation :: proc(ac: ^AnimationController, anim: ^CardSlideAnimation) {
	append(&ac.cardSlideAnimations, anim)
}

// Update processes all active animations
UpdateAnimations :: proc(ac: ^AnimationController, dt: f32) {
	// Update card attack animations
	activeAttackAnims: [dynamic]^CardAnimation
	for anim in ac.cardAttackAnimations {
		anim.Progress += dt * 2.5 // Adjust speed here
		if anim.Progress >= 1.0 {
			anim.OnFinish()
		} else {
			append(&activeAttackAnims, anim)
		}
	}
	ac.cardAttackAnimations = activeAttackAnims

	// Update card slide animations
	activeSlideAnims: [dynamic]^CardSlideAnimation
	for anim in ac.cardSlideAnimations {
		anim.Progress += dt * 5.0 // Fast, responsive sliding
		if anim.Progress >= 1.0 {
			anim.Card.position = anim.TargetPosition
		} else {
			anim.Card.position = linalg.lerp(
				anim.StartPosition,
				anim.TargetPosition,
				anim.Progress,
			)
			append(&activeSlideAnims, anim)
		}
	}
	ac.cardSlideAnimations = activeSlideAnims

	// Update UI fade animation
	if ac.uiFadeAnimation != nil {
		if ac.uiAlpha < ac.uiFadeAnimation.TargetAlpha {
			// Fading in
			ac.uiAlpha += dt * ac.uiFadeAnimation.Speed
			if ac.uiAlpha >= ac.uiFadeAnimation.TargetAlpha {
				ac.uiAlpha = ac.uiFadeAnimation.TargetAlpha
				if ac.uiFadeAnimation.OnComplete != nil {
					ac.uiFadeAnimation.OnComplete()
				}
				ac.uiFadeAnimation = nil
			}
		} else if ac.uiAlpha > ac.uiFadeAnimation.TargetAlpha {
			// Fading out
			ac.uiAlpha -= dt * ac.uiFadeAnimation.Speed
			if ac.uiAlpha <= ac.uiFadeAnimation.TargetAlpha {
				ac.uiAlpha = ac.uiFadeAnimation.TargetAlpha
				if ac.uiFadeAnimation.OnComplete != nil {
					ac.uiFadeAnimation.OnComplete()
				}
				ac.uiFadeAnimation = nil
			}
		}
	}
}

// DrawCardAttackAnimations renders all active card attack animations
DrawCardAttackAnimations :: proc(ac: ^AnimationController, camera: rl.Camera3D) {
	for anim in ac.cardAttackAnimations {
		start := anim.StartPosition
		target3D := anim.TargetEnemy.visualPos
		targetScreen := rl.GetWorldToScreen(rl.Vector3{target3D.x, 0.5, target3D.z}, camera)

		pos := linalg.lerp(start, targetScreen, anim.Progress)
		scale := rl.Lerp(1.0, 0.2, anim.Progress)

		// Draw card at interpolated position, shrinking
		tex := anim.Card.texture

		// Rotate it as it flies
		rotation := anim.Progress * 360.0 * 2.0

		rl.DrawTextureEx(tex^, pos, rotation, scale, rl.WHITE)
	}
}

// IsCardAnimating checks if a specific card is currently in a slide animation
IsCardAnimating :: proc(ac: ^AnimationController, cardID: uuid.Identifier) -> bool {
	for anim in ac.cardSlideAnimations {
		if anim.Card.id == cardID {
			return true
		}
	}
	return false
}

// FadeUITo starts a UI fade animation to the target alpha value
FadeUITo :: proc(ac: ^AnimationController, targetAlpha, speed: f32, onComplete: proc()) {
	ac.uiFadeAnimation = &UIFadeAnimation {
		TargetAlpha = targetAlpha,
		Speed = speed,
		OnComplete = onComplete,
	}
}

// GetUIAlpha returns the current UI alpha value
GetUIAlpha :: proc(ac: ^AnimationController) -> f32 {
	return ac.uiAlpha
}

// SetUIAlpha directly sets the UI alpha value (use when not animating)
SetUIAlpha :: proc(ac: ^AnimationController, alpha: f32) {
	ac.uiAlpha = alpha
}

// UpdateEnemies processes all enemy animations
UpdateEnemyAnimations :: proc(
	ac: ^AnimationController,
	dt: f32,
	enemies: [dynamic]Enemy,
	g: ^Game,
) {
	for &e in enemies {
		UpdateEnemy(&e, dt, g)
	}
}
