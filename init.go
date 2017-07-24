package main

import (
	"github.com/paulsevere/reactive/color"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Decorated, glfw.False)

	window, err := glfw.CreateWindow(width, height, "Conway's Game of Life", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	teal := color.New(0, 0.5, 0.5, 0.5).Compile()
	// teal, err := compileShader(color.New(0, 0.5, 0.5, 0.5).Render(), gl.FRAGMENT_SHADER)
	// if err != nil {
	// 	panic(err)
	// }
	// red, err := compileShader(color.New(1, 0.5, 0.5, 0.5).Render(), gl.FRAGMENT_SHADER)
	// if err != nil {
	// 	panic(err)
	// }

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, teal)

	gl.LinkProgram(prog)
	return prog
}
