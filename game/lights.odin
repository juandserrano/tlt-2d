package game

import "core:fmt"
import "core:strings"
import rl "vendor:raylib"

MaxLights :: 4

LightType :: enum {
	LightDirectional,
	LightPoint,
}

Light :: struct {
	Enabled:    bool,
	Type:       LightType,
	Position:   rl.Vector3,
	Target:     rl.Vector3,
	Color:      rl.Color,
	Intensity:  f32,

	// Shader Locations (Private handles to GPU memory)
	enabledLoc: i32,
	typeLoc:    i32,
	posLoc:     i32,
	targetLoc:  i32,
	colorLoc:   i32,
}

// CreateLight sets up a light and finds its variable locations in the shader
CreateLight :: proc(
	shader: rl.Shader,
	id: int,
	lType: LightType,
	pos, target: rl.Vector3,
	color: rl.Color,
	intensity: f32,
) -> Light {
	l := Light {
		Enabled   = true,
		Type      = lType,
		Position  = pos,
		Target    = target,
		Color     = color,
		Intensity = intensity,
	}

	// This assumes your shader uses variable names like "lights[0].position"
	shaderName := fmt.tprintf("lights[%d]", id)
	c_str := strings.clone_to_cstring(fmt.tprintf("%s.enabled", shaderName))
	defer delete(c_str)
	l.enabledLoc = rl.GetShaderLocation(shader, c_str)
	c_str = strings.clone_to_cstring(fmt.tprintf("%s.type", shaderName))
	l.typeLoc = rl.GetShaderLocation(shader, c_str)
	c_str = strings.clone_to_cstring(fmt.tprintf("%s.position", shaderName))
	l.posLoc = rl.GetShaderLocation(shader, c_str)
	c_str = strings.clone_to_cstring(fmt.tprintf("%s.target", shaderName))
	l.targetLoc = rl.GetShaderLocation(shader, c_str)
	c_str = strings.clone_to_cstring(fmt.tprintf("%s.color", shaderName))
	l.colorLoc = rl.GetShaderLocation(shader, c_str)

	UpdateLightValues(shader, l)
	return l
}

// UpdateLightValues sends the Go data to the GPU
UpdateLightValues :: proc(shader: rl.Shader, l: Light) {
	// Convert Bool to Int for Shader
	enabled := f32(0.0)
	if l.Enabled {
		enabled = 1.0
	}
	lType := i32(l.Type)

	// Convert Color to normalized Float array (0.0 - 1.0)
	color := []f32 {
		(f32(l.Color.r) / 255.0) * l.Intensity,
		(f32(l.Color.g) / 255.0) * l.Intensity,
		(f32(l.Color.b) / 255.0) * l.Intensity,
		(f32(l.Color.a) / 255.0) * l.Intensity,
	}

	// Send to Shader
	rl.SetShaderValue(shader, l.enabledLoc, &enabled, .FLOAT)
	rl.SetShaderValue(shader, l.typeLoc, &lType, .INT)
	pos := l.Position
	rl.SetShaderValue(shader, l.posLoc, &pos, .VEC3)
	target := l.Target
	rl.SetShaderValue(shader, l.targetLoc, &target, .VEC3)
	rl.SetShaderValue(shader, l.colorLoc, &color, .VEC4)

}
