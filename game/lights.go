package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const MaxLights = 4

// LightType defines if it's a Sun (Directional) or Lamp (Point)
type LightType int

const (
	LightDirectional LightType = 0
	LightPoint       LightType = 1
)

type Light struct {
	Enabled   bool
	Type      LightType
	Position  rl.Vector3
	Target    rl.Vector3
	Color     rl.Color
	Intensity float32

	// Shader Locations (Private handles to GPU memory)
	enabledLoc int32
	typeLoc    int32
	posLoc     int32
	targetLoc  int32
	colorLoc   int32
}

// CreateLight sets up a light and finds its variable locations in the shader
func CreateLight(shader rl.Shader, id int, lType LightType, pos, target rl.Vector3, color rl.Color, intensity float32) Light {
	l := Light{
		Enabled:   true,
		Type:      lType,
		Position:  pos,
		Target:    target,
		Color:     color,
		Intensity: intensity,
	}

	// This assumes your shader uses variable names like "lights[0].position"
	shaderName := fmt.Sprintf("lights[%d]", id)

	l.enabledLoc = rl.GetShaderLocation(shader, shaderName+".enabled")
	l.typeLoc = rl.GetShaderLocation(shader, shaderName+".type")
	l.posLoc = rl.GetShaderLocation(shader, shaderName+".position")
	l.targetLoc = rl.GetShaderLocation(shader, shaderName+".target")
	l.colorLoc = rl.GetShaderLocation(shader, shaderName+".color")

	UpdateLightValues(shader, l)
	return l
}

// UpdateLightValues sends the Go data to the GPU
func UpdateLightValues(shader rl.Shader, l Light) {
	// Convert Bool to Int for Shader
	enabled := float32(0.0)
	if l.Enabled {
		enabled = 1.0
	}
	lType := int32(l.Type)

	// Convert Color to normalized Float array (0.0 - 1.0)
	color := []float32{
		(float32(l.Color.R) / 255.0) * l.Intensity,
		(float32(l.Color.G) / 255.0) * l.Intensity,
		(float32(l.Color.B) / 255.0) * l.Intensity,
		(float32(l.Color.A) / 255.0) * l.Intensity,
	}

	// Send to Shader
	rl.SetShaderValue(shader, l.enabledLoc, []float32{enabled}, rl.ShaderUniformFloat)
	rl.SetShaderValue(shader, l.typeLoc, []float32{float32(lType)}, rl.ShaderUniformInt)
	rl.SetShaderValue(shader, l.posLoc, []float32{l.Position.X, l.Position.Y, l.Position.Z}, rl.ShaderUniformVec3)
	rl.SetShaderValue(shader, l.targetLoc, []float32{l.Target.X, l.Target.Y, l.Target.Z}, rl.ShaderUniformVec3)
	rl.SetShaderValue(shader, l.colorLoc, color, rl.ShaderUniformVec4)

}
