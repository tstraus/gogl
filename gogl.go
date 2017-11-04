package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"

	"./glutil"
)

const (
	w = 640
	h = 480
)

var (
	triangle = []float32 {
		-0.25, 0.5, 0, // top
		-0.75, -0.5, 0, // left
		0.25, -0.5, 0, // right
	}

	triangle2 = []float32 {
		0.25, -0.5, 0, // bottom
		-0.25, 0.5, 0, // left
		0.75, 0.5, 0, // right
	}

	vertexShaderSource = `
		#version 410
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 410
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(1, 1, 1, 1);
		}
	` + "\x00"
)

func draw(objects []uint32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	for _, object := range objects {
		gl.BindVertexArray(object)
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle) / 3))
	}

	window.SwapBuffers()
	glfw.PollEvents()
}

func main() {
	window, program := glutil.InitGL(w, h, vertexShaderSource, fragmentShaderSource)
	defer glfw.Terminate() // wait for init

	gl.ClearColor(0, 0.5, 1.0, 1.0)

	var objects []uint32
	objects = append(objects, glutil.MakeVao(triangle))
	objects = append(objects, glutil.MakeVao(triangle2))

	for !window.ShouldClose() {
		draw(objects, window, program)
	}
}

// ubuntu needed: sudo apt install libglu1-mesa-dev -y
