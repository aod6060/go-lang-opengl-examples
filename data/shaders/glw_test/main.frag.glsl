#version 410

uniform sampler2D tex0;

in vec2 v_TexCoords;
in vec3 v_Normal;

out vec4 out_Color;

void main() {
    out_Color = texture(tex0, v_TexCoords);
}