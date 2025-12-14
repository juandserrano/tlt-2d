#version 330

in vec3 fragPosition;
in vec3 fragNormal;
in vec2 fragTexCoord;
in vec4 fragColor;

out vec4 finalColor;

// --- STANDARD RAYLIB INPUTS ---
uniform sampler2D texture0;
uniform vec4 colDiffuse;

// --- THE LIGHTING DATA STRUCTURE ---
// This matches the Go struct we will create below
struct Light {
    int enabled;
    int type; // 0=Directorional, 1=Point
    vec3 position;
    vec3 target;
    vec4 color;
};

// We allow up to 4 lights
uniform Light lights[4];
uniform vec4 ambient;

void main() {
    vec4 texelColor = texture(texture0, fragTexCoord);
    vec3 lightDot = vec3(0.0);
    vec3 normal = normalize(fragNormal);

    // Loop through all 4 lights
    for (int i = 0; i < 4; i++) {
        if (lights[i].enabled == 1) {
            vec3 lightDir;
            
            // Logic for Point Light vs Directional Light
            if (lights[i].type == 0) {
                // Directional: Like the sun, coming from specific angle
                lightDir = -normalize(lights[i].target - lights[i].position);
            } else {
                // Point: Like a lightbulb, coming from a specific spot
                lightDir = normalize(lights[i].position - fragPosition);
            }

            // Calculate Brightness (Dot Product)
            float NdotL = max(dot(normal, lightDir), 0.0);
            
            // Add this light's contribution to the total
            lightDot += lights[i].color.rgb * NdotL;
        }
    }
    // DEBUG 1: Visualizing Normals
    // If your object turns colorful (Blue/Green/Red gradients), Normals are working!
    // If it is black, the Vertex Shader is not sending data.
    finalColor = vec4(normalize(fragNormal) * 0.5 + 0.5, 1.0); return;

    // DEBUG 2: Check Distance
    // Make everything White. If you see the object, geometry is fine.
    // finalColor = vec4(1.0, 1.0, 1.0, 1.0); return;
    
    // Final Math: (Ambient + Sum of all Lights) * Texture
    vec3 finalRGB = (ambient.rgb + lightDot) * texelColor.rgb * colDiffuse.rgb;
    finalColor = vec4(finalRGB, texelColor.a);
}