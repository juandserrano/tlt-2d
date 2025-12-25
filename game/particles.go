package game

import (
	"image/color"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Particle struct {
	Position     rl.Vector3
	Velocity     rl.Vector3
	Acceleration rl.Vector3
	Color        color.RGBA
	StartColor   color.RGBA
	EndColor     color.RGBA
	Size         float32
	StartSize    float32
	EndSize      float32
	Life         float32
	MaxLife      float32
	Rotation     float32
	RotSpeed     float32
	Texture      *rl.Texture2D
	Active       bool
}

type ParticleManager struct {
	particles      []*Particle
	defaultTexture rl.Texture2D
	useDefaultText bool
	maxParticles   int
}

func NewParticleManager(maxParticles int) *ParticleManager {
	pm := &ParticleManager{
		maxParticles: maxParticles,
		particles:    make([]*Particle, maxParticles),
	}

	for i := 0; i < maxParticles; i++ {
		pm.particles[i] = &Particle{Active: false}
	}

	return pm
}

func (pm *ParticleManager) Init() {
	// Generate a default soft circle texture
	// Gradient radial: center white, edge transparent
	img := rl.GenImageGradientRadial(32, 32, 0.0, rl.White, rl.Blank)
	pm.defaultTexture = rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	pm.useDefaultText = true
}

func (pm *ParticleManager) Unload() {
	if pm.useDefaultText {
		rl.UnloadTexture(pm.defaultTexture)
	}
}

func (pm *ParticleManager) Update(dt float32) {
	for _, p := range pm.particles {
		if !p.Active {
			continue
		}

		p.Life -= dt
		if p.Life <= 0 {
			p.Active = false
			continue
		}

		// Physics
		p.Velocity = rl.Vector3Add(p.Velocity, rl.Vector3Scale(p.Acceleration, dt))
		p.Position = rl.Vector3Add(p.Position, rl.Vector3Scale(p.Velocity, dt))
		p.Rotation += p.RotSpeed * dt

		// Interpolation
		t := 1.0 - (p.Life / p.MaxLife)
		p.Size = rl.Lerp(p.StartSize, p.EndSize, t)
		p.Color = LerpColor(p.StartColor, p.EndColor, t)
	}
}

func (pm *ParticleManager) Draw(camera rl.Camera3D) {
	rl.BeginBlendMode(rl.BlendAlpha)
	// Disable depth mask so particles don't occlude each other weirdly if they are transparent
	// rl.BeginMode3D should already be active
	// rl.DrawBillboard uses the camera to align the quad

	for _, p := range pm.particles {
		if !p.Active {
			continue
		}

		tex := pm.defaultTexture
		if p.Texture != nil {
			tex = *p.Texture
		}

		// Draw billboards
		// Note: Raylib-go DrawBillboard might not support rotation directly without Pro.
		// Using standard DrawBillboard for now.
		rl.DrawBillboard(camera, tex, p.Position, p.Size, p.Color)
	}
	rl.EndBlendMode()
}

func (pm *ParticleManager) Spawn(p Particle) {
	// Find first inactive slot
	for _, slot := range pm.particles {
		if !slot.Active {
			*slot = p
			slot.Active = true
			return
		}
	}
}

func (pm *ParticleManager) SpawnExplosion(pos rl.Vector3, count int, col color.RGBA) {
	for i := 0; i < count; i++ {
		// Random spread
		spread := float32(2.0)
		vel := rl.Vector3{
			X: (rand.Float32() - 0.5) * spread * 2,
			Y: (rand.Float32() - 0.5) * spread * 2 + 3.0, // Upward burst
			Z: (rand.Float32() - 0.5) * spread * 2,
		}

		p := Particle{
			Position:     pos,
			Velocity:     vel,
			Acceleration: rl.Vector3{0, -5.0, 0}, // Gravity
			StartColor:   col,
			EndColor:     color.RGBA{col.R, col.G, col.B, 0},
			StartSize:    0.3 + rand.Float32()*0.2,
			EndSize:      0.0,
			Life:         0.5 + rand.Float32()*0.5,
			MaxLife:      1.0, // adjusted locally
			Active:       true,
		}
		p.MaxLife = p.Life
		pm.Spawn(p)
	}
}

func LerpColor(start, end color.RGBA, t float32) color.RGBA {
	return color.RGBA{
		R: uint8(float32(start.R) + t*(float32(end.R)-float32(start.R))),
		G: uint8(float32(start.G) + t*(float32(end.G)-float32(start.G))),
		B: uint8(float32(start.B) + t*(float32(end.B)-float32(start.B))),
		A: uint8(float32(start.A) + t*(float32(end.A)-float32(start.A))),
	}
}
