package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, x float64, z float64) {
	reader := bufio.NewReader(strings.NewReader(s))
	x0, z0 := drive(reader)
	if math.Abs(x0-x) > 1e-6 || math.Abs(z0-z) > 1e-6 {
		t.Errorf("Sample expect %f, %f, but got %f, %f", x, z, x0, z0)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `7 2 11
3 -11 2
`, 6.5, 2.0)
}

func TestSample2(t *testing.T) {
	runSample(t, `7 2 11
4 -3 3
`, 4.1666666667, 1.0000000000)
}
