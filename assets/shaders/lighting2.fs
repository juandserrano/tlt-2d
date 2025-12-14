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
    float enabled;
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
        if (lights[i].enabled == 1.0) {
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
    // finalColor = vec4(normalize(fragNormal) * 0.5 + 0.5, 1.0); return;

    // DEBUG 2: Check Distance
    // Make everything White. If you see the object, geometry is fine.
    // finalColor = vec4(1.0, 1.0, 1.0, 1.0); return;
    
    // Final Math: (Ambient + Sum of all Lights) * Texture
    vec3 finalRGB = (ambient.rgb + lightDot) * texelColor.rgb * colDiffuse.rgb;
    finalColor = vec4(finalRGB, texelColor.a);
}



// TESTS

// #version 330

// in vec3 fragPosition;
// in vec3 fragNormal;
// in vec2 fragTexCoord;
// in vec4 fragColor;

// out vec4 finalColor;

// uniform sampler2D texture0;
// uniform vec4 colDiffuse;
// uniform vec4 ambient;

// struct Light {
//     float enabled;
//     int type;
//     vec3 position;
//     vec3 target;
//     vec4 color;
// };

// uniform Light lights[4];

// void main() {
//     // TEST 1: Is the Texture/Material Color Black?
//     // If the object is pure RED, this part is fine.
//     // If Black, your colDiffuse or Texture is broken.
//     // finalColor = vec4(1.0, 0.0, 0.0, 1.0) * colDiffuse * texture(texture0, fragTexCoord); return;

//     // TEST 2: Is Light 0 Enabled?
//     // If the object is GREEN, the shader knows Light 0 is enabled.
//     // If Black, the 'enabled' integer is not reaching the GPU.
//     if (lights[0].enabled > 0.5) {
//          finalColor = vec4(0.0, 1.0, 0.0, 1.0); return;
//     }

//     // TEST 3: Is Light 0 Color reaching the GPU?
//     // If the object matches your light color (e.g. Red), the color data is correct.
//     // If Black, UpdateLightValues in Go is failing.
//     // finalColor = lights[0].color; return;

//     // TEST 4: The Lighting Math (NdotL)
//     // We manually calculate Light 0.
//     // If you see gradients (shading), the math works. 
//     // If you see solid black, the Light Position is wrong (or inside the object).
//     // vec3 lightDir = normalize(lights[0].position - fragPosition);
//     // float NdotL = max(dot(normalize(fragNormal), lightDir), 0.0);
//     // finalColor = vec4(vec3(NdotL), 1.0); return;
// }