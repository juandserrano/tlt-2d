#version 330

// Input vertex attributes (from vertex shader)
in vec3 fragPosition;
in vec2 fragTexCoord;
in vec4 fragColor;
in vec3 fragNormal;

// Input uniform values
uniform sampler2D texture0;
uniform vec4 colDiffuse;

// Output fragment color
out vec4 finalColor;

// Custom uniforms
uniform vec3 lightPos;
uniform vec4 ambient;
uniform vec3 viewPos;

void main()
{
    // Texture color
    vec4 texelColor = texture(texture0, fragTexCoord);

    // Ambient light
    vec3 ambientLight = ambient.rgb;

    // Diffuse light
    vec3 norm = normalize(fragNormal);
    vec3 lightDir = normalize(lightPos - fragPosition);
    float diff = max(dot(norm, lightDir), 0.0);
    vec3 diffuse = diff * vec3(1.0, 1.0, 1.0); // White light

    // Specular light (optional, simple Blinn-Phong)
    float specularStrength = 0.5;
    vec3 viewDir = normalize(viewPos - fragPosition);
    vec3 reflectDir = reflect(-lightDir, norm);  
    float spec = pow(max(dot(viewDir, reflectDir), 0.0), 32);
    vec3 specular = specularStrength * spec * vec3(1.0, 1.0, 1.0);

    // Final color
    vec3 result = (ambientLight + diffuse + specular) * texelColor.rgb;
    finalColor = vec4(result, texelColor.a);
}
