#version 330
in vec3 fragNormal;
in vec3 fragPosition;
uniform vec4 colDiffuse;
out vec4 finalColor;
void main() {
    finalColor = colDiffuse;
}
