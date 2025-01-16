#version 410

uniform sampler2D tex0;

in vec4 v_colors;
in vec2 v_TexCoords;

out vec4 out_Color;

void main() {
    out_Color = v_colors * texture(tex0, v_TexCoords);
}