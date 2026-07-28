package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect float64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if math.Abs(res-expect) > 1e-6 && math.Abs(res-expect)/math.Max(1, math.Abs(expect)) > 1e-6 {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2
2 0
0 2
`, 90)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
2 0
0 2
-2 2
`, 135)
}

func TestSample3(t *testing.T) {
	runSample(t, `4
2 0
0 2
-2 0
0 -2
`, 270)
}

func TestSample4(t *testing.T) {
	runSample(t, `2
2 1
1 2
`, 36.8698976458)
}
