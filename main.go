package main

import (

	// . "./point"
	"runtime"

	"github.com/davecgh/go-spew/spew"
	"github.com/paulsevere/reactive/element"
	// . "./color"
	// . "github.com/davecgh/go-spew/spew"
	"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width  = 1000
	height = 1000
)

func main() {
	spew.Dump(gl.COLOR_ATTACHMENT2)

	runtime.LockOSThread()
	box := element.New(300, 100, 600, 0)
	box2 := element.New(400, 400, 100, 0)
	// box2.Glp = box.ToMat()

	// spew.Dump(box, box2)

	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	// vao := makeVao(points)

	shapes := element.NewEls()
	shapes.AddElement(box)
	shapes.AddElement(box2)

	// spew.Dump(shapes)

	for !window.ShouldClose() {
		draw(shapes.Nodes, window, program)

	}
}

func draw(shapes []uint32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)
	gl.ClearColor(0, 0.35, 0.5, 1)

	for _, ob := range shapes {
		gl.BindVertexArray(ob)
		gl.DrawArrays(gl.TRIANGLES, 0, 6)
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
