#version 330

// Inputs
in vec2 fragTexCoord;
in vec4 fragColor;

// Uniforms
uniform sampler2D texture0;
uniform vec4 colDiffuse;
uniform vec2 textureSize;
uniform float time;

uniform vec3 glowColor;
uniform float radius;

out vec4 finalColor;

void main() {
    // 1. Get the original color
    vec4 texel = texture(texture0, fragTexCoord);

    // If the pixel is already transparent, we don't draw anything (it's empty space)
    if (texel.a == 0.0) discard;

    // 2. Inner Glow Parameters
    // float radius = 6.0; // The "6px inwards" you requested
    // vec3 glowColor = vec3(0.0, 0.8, 1.0); // Cyan Blue
    float safeRadius = clamp(radius, 0.0, 20.0);

    // 3. Scan for the nearest "Edge"
    // We start assuming the edge is far away (at distance = radius)
    float minDistanceToEdge = safeRadius;
    
    vec2 size = 1.0 / textureSize;

    // Loop through the box [-radius, +radius]
    for (float x = -safeRadius; x <= safeRadius; x++) {
        for (float y = -safeRadius; y <= safeRadius; y++) {
            
            // Optimization: Skip checking center
            if (x == 0.0 && y == 0.0) continue;

            vec2 offset = vec2(x, y) * size;
            vec2 neighborUV = fragTexCoord + offset;

            // Check if this neighbor is "The Edge"
            // It is an edge if:
            // A. The neighbor is outside the texture bounds (0.0 to 1.0)
            // B. The neighbor pixel is transparent
            bool isEdge = false;
            
            if (neighborUV.x < 0.0 || neighborUV.x > 1.0 || 
                neighborUV.y < 0.0 || neighborUV.y > 1.0) {
                isEdge = true;
            } else {
                float alpha = texture(texture0, neighborUV).a;
                if (alpha == 0.0) isEdge = true;
            }

            // If we found an edge, measure how close it is
            if (isEdge) {
                float dist = length(vec2(x, y));
                if (dist < minDistanceToEdge) {
                    minDistanceToEdge = dist;
                }
            }
        }
    }

    // 4. Calculate Intensity
    // If minDistanceToEdge is 0 (we are at the edge), intensity is 1.0
    // If minDistanceToEdge is 6 (we are deep inside), intensity is 0.0
    float intensity = 1.0 - (minDistanceToEdge / safeRadius);
    
    // Clamp to 0 just in case
    intensity = max(intensity, 0.0);

    // 5. Apply Pulsating Logic
    float pulse = (sin(time * 4.0) * 0.4) + 0.6; // Pulse between 0.2 and 1.0
    intensity *= pulse;

    // 6. Mix the Glow ON TOP of the original texture
    // mix(Original, GlowColor, Factor)
    vec3 mixedRGB = mix(texel.rgb, glowColor, intensity);

    finalColor = vec4(mixedRGB, texel.a);
}