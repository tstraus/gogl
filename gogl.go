package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"

	"./glshapes"
	"./glutil"
)

const (
	w = 1280
	h = 720
)

var (
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
			frag_colour = vec4(1, 0.5, 0, 1);
		}
	` + "\x00"
)

func draw(objects []*glshapes.Shape, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	for _, object := range objects {
		gl.BindVertexArray(object.Vao)
		gl.DrawArrays(gl.TRIANGLES, 0, object.PolygonCount)
	}

	window.SwapBuffers()
	glfw.PollEvents()
}

func main() {
	window, program := glutil.InitGL(w, h, vertexShaderSource, fragmentShaderSource)
	defer glfw.Terminate() // wait for init

	gl.ClearColor(0, 0.5, 1.0, 1.0)

	var objects []*glshapes.Shape
	objects = append(objects, glshapes.NewRectangle(glshapes.Vec2f{-0.5, -0.5}, glshapes.Vec2f{0.5, 0.5}))

	for !window.ShouldClose() {
		draw(objects, window, program)
	}
}

// ubuntu needed: sudo apt install libglu1-mesa-dev -y
