package p3594

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, k int, m int, time []int, mul []float64, expect float64) {
	res := minTime(n, k, m, time, mul)
	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	k := 1
	m := 2
	time := []int{5}
	mul := []float64{1.0, 1.3}
	expect := 5.0
	runSample(t, n, k, m, time, mul, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	k := 2
	m := 3
	time := []int{2, 5, 8}
	mul := []float64{1.0, 1.5, 0.75}
	expect := 14.5
	runSample(t, n, k, m, time, mul, expect)
}

func TestSample3(t *testing.T) {
	n := 1
	k := 4
	m := 5
	time := []int{58}
	mul := []float64{0.96, 0.69, 0.87, 1.97, 0.7}
	expect := 55.68
	runSample(t, n, k, m, time, mul, expect)
}

func TestSample4(t *testing.T) {
	n := 2
	k := 3
	m := 4
	time := []int{40, 1}
	mul := []float64{1.82, 1.59, 1.11, 1.84}
	expect := 47.81
	runSample(t, n, k, m, time, mul, expect)
}

func TestSample5(t *testing.T) {
	n := 4
	k := 2
	m := 1
	time := []int{51, 34, 25, 18}
	mul := []float64{1.23}
	expect := 177.12
	runSample(t, n, k, m, time, mul, expect)
}
