#version 330
in vec3 vertexPosition;
in vec3 vertexNormal;
uniform mat4 mvp;
out vec3 fragNormal;
out vec3 fragPosition;
void main() {
    fragPosition = vertexPosition;
    fragNormal = vertexNormal;
    gl_Position = mvp * vec4(vertexPosition, 1.0);
}
