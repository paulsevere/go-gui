package color

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Color struct {
	R       float32
	G       float32
	B       float32
	Opacity float32
}

func New(R float32,
	G float32,
	B float32, O float32) Color {
	return Color{R: R, G: G, B: B, Opacity: O}
}

func (col Color) Render() string {
	return fmt.Sprintf(`
		#version 410
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(%g, %g, %g, %g);
		}
	`, col.R, col.G, col.B, col.Opacity) + "\x00"
}

func (col Color) Compile() uint32 {
	marker, _ := compileShader(col.Render(), gl.FRAGMENT_SHADER)
	return marker
}

func Red() string {
	return `
		#version 410
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(1, 0, 0, 1.0);
		}
	` + "\x00"
}

func Blue() Color {
	return New(0, 0, 1, 1)
}
