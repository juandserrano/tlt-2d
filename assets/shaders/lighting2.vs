#version 330

// --- 1. INPUTS (Data from your .obj mesh) ---
in vec3 vertexPosition;
in vec2 vertexTexCoord;
in vec3 vertexNormal;
in vec4 vertexColor;

// --- 2. UNIFORMS (Matrices sent by Raylib) ---
uniform mat4 mvp;       // Camera + Object Transform (Model View Projection)
uniform mat4 matModel;  // Object Transform only (Position/Rotation in world)
uniform mat4 matNormal; // Rotation matrix specifically for normals

// --- 3. OUTPUTS (Sent to Fragment Shader) ---
out vec3 fragPosition;
out vec2 fragTexCoord;
out vec3 fragNormal;
out vec4 fragColor;

void main()
{
    // A. Calculate Position on Screen (MANDATORY for all vertex shaders)
    gl_Position = mvp * vec4(vertexPosition, 1.0);

    // B. Calculate Position in World Space (Needed for lighting distance)
    fragPosition = vec3(matModel * vec4(vertexPosition, 1.0));

    // C. Pass Texture Coordinates and Color
    fragTexCoord = vertexTexCoord;
    fragColor = vertexColor;

    // D. Calculate Normal in World Space (The part causing your error)
    // We take the mesh's normal and rotate it by the object's rotation (matNormal).
    // This ensures that if you rotate the cube, the light still bounces off the correct side.
    fragNormal = normalize(vec3(matNormal * vec4(vertexNormal, 1.0)));
}