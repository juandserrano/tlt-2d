#version 330

in vec2 fragTexCoord;
in vec3 fragNormal;
in vec3 fragPosition;

out vec4 finalColor;

uniform vec3 viewPos; // Camera position

void main() {
    // 1. Basic Colors
    vec3 waterColor = vec3(0.0, 0.4, 0.8); // Deep Blue
    vec3 foamColor = vec3(1.0, 1.0, 1.0);  // White

    // 2. Lighting Setup
    vec3 lightDir = normalize(vec3(0.5, 1.0, -0.5)); // Sun position
    vec3 viewDir = normalize(viewPos - fragPosition);
    vec3 normal = normalize(fragNormal);

    // 3. Specular Highlight (Sun reflection)
    vec3 reflectDir = reflect(-lightDir, normal);
    float spec = pow(max(dot(viewDir, reflectDir), 0.0), 32.0);

    // 4. Height-based coloring (lighter at tips, darker at bottom)
    // We use the Y height to mix foam
    float heightFactor = smoothstep(-1.0, 1.5, fragPosition.y);
    vec3 col = mix(waterColor * 0.5, waterColor, heightFactor);

    // Add specular
    col += vec3(0.3) * spec;

    finalColor = vec4(col, 0.8); // 0.8 alpha for slight transparency
}