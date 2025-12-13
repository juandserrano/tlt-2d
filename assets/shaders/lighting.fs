#version 330

// Input vertex attributes (from vertex shader)
in vec3 fragPosition;
in vec2 fragTexCoord;
in vec4 fragColor;
in vec3 fragNormal;
in vec4 fragPosLightSpace;

// Input uniform values
uniform sampler2D texture0;
uniform vec4 colDiffuse;

// Output fragment color
out vec4 finalColor;

// Custom uniforms
uniform vec3 lightPos;
uniform vec4 ambient;
uniform vec3 viewPos;
uniform sampler2D texture_emission;
uniform sampler2D shadowMap;

float ShadowCalculation(vec4 fragPosLightSpace)
{
    // perform perspective divide
    vec3 projCoords = fragPosLightSpace.xyz / fragPosLightSpace.w;
    // transform to [0,1] range
    projCoords = projCoords * 0.5 + 0.5;
    
    // get closest depth value from light's perspective (using [0,1] range fragPosLight as coords)
    float closestDepth = texture(shadowMap, projCoords.xy).r; 
    
    // get depth of current fragment from light's perspective
    float currentDepth = projCoords.z;
    
    // check whether current frag pos is in shadow
    // float bias = 0.005;
    // float bias = max(0.05 * (1.0 - dot(normal, lightDir)), 0.005);
    float bias = 0.002; 
    float shadow = currentDepth - bias > closestDepth ? 1.0 : 0.0;

    if(projCoords.z > 1.0)
        shadow = 0.0;
        
    return shadow;
}

void main()
{
    // Texture color
    vec4 texelColor = texture(texture0, fragTexCoord);
    vec4 emissionColor = texture(texture_emission, fragTexCoord);

    // Ambient light
    vec3 ambientLight = ambient.rgb;

    // Diffuse light
    vec3 norm = normalize(fragNormal);
    vec3 lightDir = normalize(lightPos);
    float diff = max(dot(norm, lightDir), 0.0);
    vec3 diffuse = diff * vec3(1.0, 1.0, 1.0); // White light

    // Shadow calculation
    float shadow = ShadowCalculation(fragPosLightSpace);

    // Final color
    vec3 result = (ambientLight + (1.0 - shadow) * diffuse) * texelColor.rgb + emissionColor.rgb;
    finalColor = vec4(result, texelColor.a);
}
