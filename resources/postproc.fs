#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;
in vec4 fragColor;

// Input uniform values
uniform sampler2D texture0;
uniform vec4 colDiffuse;

// Output fragment color
out vec4 finalColor;

float renderWidth = 1280;
float renderHeight = 720;
float offset = 0.0;

uniform float time;

uniform float amount;
uniform sampler2D tDiffuse;
varying vec2 vUv;

float random(vec2 p)
{
  vec2 K1 = vec2(
    23.14069263277926, // e^pi (Gelfond's constant)
    2.665144142690225 // 2^sqrt(2) (Gelfond–Schneider constant)
  );

  return fract(cos(dot(p,K1)) * 12345.6789);
}

void main()
{
    float frequency = renderHeight/3.0;

    float globalPos = (fragTexCoord.y + offset) * frequency;
    float wavePos = cos((fract(globalPos) - 0.5)*3.14);

    // Texel color fetching from texture sampler
    vec4 texelColor = texture(texture0, fragTexCoord);

    finalColor = mix(vec4(0, 0, 0, 0.0), texelColor, wavePos);

    vec4 noiseColor = texture(texture0, fragTexCoord);
    vec2 uvRandom = fragTexCoord;
    uvRandom.y *= random(vec2(uvRandom.y,amount));
    noiseColor.rgb += random(uvRandom)*0.175;

    finalColor = mix(vec4(noiseColor), finalColor, 0.35);
}

