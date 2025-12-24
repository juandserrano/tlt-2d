#version 330

// Inputs
in vec2 fragTexCoord;
in vec4 fragColor;

// Uniforms
uniform sampler2D texture0;
uniform vec4 colDiffuse;

// Custom Uniforms (We send these from Go)
uniform vec2 textureSize; // The width and height of the texture (e.g., 64, 64)
uniform vec4 outlineColor; // The color of the line (e.g., Red)
uniform float time; // Time for pulsation

out vec4 finalColor;

void main() {
    // 1. Get the color of the pixel itself
    vec4 texel = texture(texture0, fragTexCoord);

    // 2. If the pixel is transparent, skip it (inner outline only draws on opaque pixels)
    if (texel.a <= 0.0) {
        finalColor = vec4(0.0);
        return;
    }

    // 3. We are on an opaque pixel. Check neighbors to see if we are near a transparent edge.
    // We calculate the size of "one pixel" in UV space (0.0 to 1.0)
    float xStep = 1.0 / textureSize.x;
    float yStep = 1.0 / textureSize.y;

    bool isOutline = false;

    // Check neighbors within 3 pixels
    for (int y = -3; y <= 3; y++) {
        for (int x = -3; x <= 3; x++) {
            if (x == 0 && y == 0) continue;

            vec2 offset = vec2(float(x) * xStep, float(y) * yStep);
            vec2 nearby = fragTexCoord + offset;

            // If the neighbor is outside the UV bounds (0-1), consider it transparent
            // This ensures the outline is drawn at the very edge of the texture quad
            if (nearby.x < 0.0 || nearby.x > 1.0 || nearby.y < 0.0 || nearby.y > 1.0) {
                isOutline = true;
                break;
            }

            float alpha = texture(texture0, nearby).a;

            // If we find a transparent neighbor, this texel is part of the inner border
            if (alpha <= 0.0) {
                isOutline = true;
                break;
            }
        }
        if (isOutline) break;
    }

    // 4. If it's an outline, use the outline color. Otherwise, use texture color.
    if (isOutline) {
        // Calculate pulsation (0.0 to 1.0)
        float pulse = (sin(time * 5.0) + 1.0) * 0.5;
        // Mix between original texture color and outline color
        finalColor = mix(texel * colDiffuse * fragColor, outlineColor, pulse);
    } else {
        finalColor = texel * colDiffuse * fragColor;
    }
}