#version 330

// Input vertex attributes (from Raylib)
in vec3 vertexPosition;
in vec2 vertexTexCoord;
in vec3 vertexNormal;

// Output to Fragment Shader
out vec2 fragTexCoord;
out vec3 fragNormal;
out vec3 fragPosition;

// Uniforms
uniform mat4 mvp;   // Model-View-Projection matrix
uniform mat4 matModel; // Model matrix
uniform float time;

void main() {
    vec3 pos = vertexPosition;

    // --- Wave Logic ---
    // We combine multiple sine waves to make it look irregular
    float freq = 10.0;
    float amp = 0.01;
    float speed = 5.0;
    
    // Wave 1
    pos.y += sin(pos.x * freq + time * speed) * amp;
    // Wave 2 (different direction)
    pos.y += sin(pos.z * freq * 1.5 + time * speed * 1.3) * amp * 0.5;
    // Wave 3 (diagonal)
    pos.y += sin((pos.x + pos.z) * freq * 0.5 + time * speed * 0.8) * amp * 2.0;

    // Calculate generic normal (simplified for performance)
    // In a production shader, you would calculate exact normals based on the derivative of the waves above
    vec3 newNormal = normalize(vec3(0.0, 1.0, 0.0) + pos * 0.1); 

    // Send data to fragment shader
    fragTexCoord = vertexTexCoord;
    fragNormal = newNormal;
    fragPosition = vec3(matModel * vec4(pos, 1.0));

    // Final position on screen
    gl_Position = mvp * vec4(pos, 1.0);
}