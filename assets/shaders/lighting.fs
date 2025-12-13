#version 330

// Input from Vertex Shader (standard Raylib pipeline)
in vec2 fragTexCoord;
in vec4 fragColor;

// Output to the screen
out vec4 finalColor;

// Uniforms (Data sent from Go)
uniform sampler2D texture0; // The Material's texture (sent automatically by Raylib)
uniform vec4 colDiffuse;    // The Material's tint color (sent automatically)
uniform vec4 ambientLight;  // THE CUSTOM UNIFORM we will send manually

void main()
{
    // 1. Get the pixel color from the texture at this coordinate
    vec4 texelColor = texture(texture0, fragTexCoord);

    // 2. Multiply texture color * material tint * ambient light
    // This creates the final look. If ambientLight is dark (0.2, 0.2, 0.2), the object looks dark.
    finalColor = texelColor * colDiffuse * ambientLight;
}