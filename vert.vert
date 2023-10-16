	#version 400

  in vec3 vp;
	in vec3 color;
	out vec3 vert_color;
	void main() {
		gl_Position = vec4(vp, 1.0);
		vert_color = color;
	}