package shapes

import . "../point"

type Shape interface {
	Vec() []float32
}

type Triangle struct {
	A Point
	B Point
	C Point
}

func NewTriangle(a Point, b Point, c Point) Triangle {
	return Triangle{A: a, B: b, C: c}
}

func (t Triangle) Vec() []float32 {
	points := []float32{t.A.X, t.A.Y, t.A.Z, t.B.X, t.B.Y, t.B.Z, t.C.X, t.C.Y, t.C.Z}
	return points
}
