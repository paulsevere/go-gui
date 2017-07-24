package point

type Point struct {
	X float64
	Y float64
	Z float64
}

func New(X float64, Y float64) Point {
	return Point{X: X, Y: Y, Z: 0.0}
}
