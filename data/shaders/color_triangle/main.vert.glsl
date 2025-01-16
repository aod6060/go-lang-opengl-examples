#version 410


layout(location = 0) in vec3 vertices;
layout(location = 1) in vec4 colors;


uniform mat4 proj;
uniform mat4 view;
uniform mat4 model;

out vec4 v_colors;

void main() {
    gl_Position = proj * view * model * vec4(vertices, 1.0);
    v_colors = colors;
}