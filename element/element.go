package element

import (
	"github.com/paulsevere/reactive/color"
	"github.com/paulsevere/reactive/point"
)

type Element struct {
	Width  float64
	Height float64
	Pos    point.Point
	Color  color.Color
	ctx    Ctx
	Glp    uint32
}

func New(w float64, h float64, x float64, y float64) Element {

	el := Element{Width: w, Height: h, Pos: point.New(x, y), Color: color.Blue(), ctx: NewCtx(1000, 1000), Glp: 0}
	return el
}

// func (el Element) Draw(window *glfw.Window, program uint32) {
// 	gl.UseProgram(program)

// 	for _, ob := range shapes {
// 		gl.BindVertexArray(ob)
// 		gl.DrawArrays(gl.TRIANGLES, 0, 6)
// 	}
// }

func (el Element) ToMat() uint32 {
	wid := float32(el.ctx.LenX(el.Width))
	het := float32(el.ctx.LenY(el.Height))
	x, y := el.ctx.Point(el.Pos.X, el.Pos.Y)
	orx := float32(x)
	ory := float32(y)

	all := []float32{orx, ory, 0.0, orx + wid, ory, 0.0, orx + wid, ory - het, 0.0, orx + wid, ory - het, 0.0, orx, ory - het, 0.0, orx, ory, 0.0}

	return makeVao(all)
}

type Ctx struct {
	Width  float64
	Height float64
}

func NewCtx(x float64, y float64) Ctx {
	return Ctx{Width: x, Height: y}
}

func (c Ctx) Point(x float64, y float64) (float64, float64) {
	return c.ScaleX(x), c.ScaleY(y)
}

func (c Ctx) LenX(v float64) float64 {
	return v / c.Width
}
func (c Ctx) LenY(v float64) float64 {
	return v / c.Height
}
func (c Ctx) ScaleX(v float64) float64 {
	w := c.Width
	adj := v / w
	if adj < 0.5 {
		return (-(0.5 - adj)) * 2
	} else {
		return (adj - 0.5) * 2
	}

}
func (c Ctx) ScaleY(v float64) float64 {
	h := c.Height
	adj := v / h
	if adj > 0.5 {
		return (0.5 - adj) * 2
	} else {
		return -(adj - 0.5) * 2
	}
}

func AddEl(w float64, h float64, x float64, y float64, ctx Ctx) uint32 {
	wid := float32(ctx.LenX(w))
	het := float32(ctx.LenY(h))
	dx, dy := ctx.Point(x, y)
	orx := float32(dx)
	ory := float32(dy)
	// spew.Dump(orx, ory)

	all := []float32{orx, ory, 0.0, orx + wid, ory, 0.0, orx + wid, ory - het, 0.0, orx + wid, ory - het, 0.0, orx, ory - het, 0.0, orx, ory, 0.0}

	return makeVao(all)
}
