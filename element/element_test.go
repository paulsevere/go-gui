package element

import (
	"fmt"
	"math"
	"testing"
)

func testRun(val float64, exp float64, t *testing.T) {
	ctx := NewCtx(1000, 1000)
	res := ctx.ScaleX(float64(val))

	if Round(res*100) != Round(exp*100) {
		fmt.Printf("expected %.2f -- got %.2f\n", exp, res)
		t.Fail()
	}
}
func testY(val float64, exp float64, t *testing.T) {
	ctx := NewCtx(1000, 1000)
	res := ctx.ScaleY(float64(val))

	if Round(res*100) != Round(exp*100) {
		fmt.Printf("Case: %v expected %.2f -- got %.2f\n", val, exp, res)
		t.Fail()
	}
}

func Test1(t *testing.T) {
	testRun(250, -0.5, t)
	testRun(100, -0.8, t)
	testRun(500, 0.0, t)
	testRun(900, 0.8, t)
	testY(250, 0.5, t)
	testY(100, 0.8, t)
	testY(500, 0.0, t)
	testY(900, -0.8, t)
	testY(0, 1, t)

}

func Round(f float64) float64 {
	return math.Floor(f + .5)
}

func RoundPlus(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return Round(f*shift) / shift
}
