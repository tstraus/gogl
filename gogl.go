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

func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	window.SwapBuffers()
	glfw.PollEvents()
}

func main() {
	window, program := glutil.InitGL(w, h)
	defer glfw.Terminate() // wait for init

	gl.ClearColor(0, 0.5, 1.0, 1.0)

	for !window.ShouldClose() {
		draw(window, program)
	}
}

// ubuntu needed: sudo apt install libglu1-mesa-dev -y
