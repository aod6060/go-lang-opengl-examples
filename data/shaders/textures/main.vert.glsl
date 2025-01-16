#version 410


layout(location = 0) in vec3 vertices;
layout(location = 1) in vec4 colors;
layout(location = 2) in vec2 texCoords;

uniform mat4 proj;
uniform mat4 view;
uniform mat4 model;

out vec4 v_colors;
out vec2 v_TexCoords;

void main() {
    gl_Position = proj * view * model * vec4(vertices, 1.0);
    v_colors = colors;
    v_TexCoords = texCoords;
}