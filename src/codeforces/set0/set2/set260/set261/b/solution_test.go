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
	if math.Abs(res-expect) > 1e-4 && math.Abs(res-expect)/math.Max(1, math.Abs(expect)) > 1e-4 {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
1 2 3
3
`, 1.3333333333)
}

func TestSample2(t *testing.T) {
	runSample(t, `9
2 2 2 2 2 2 2 1 2
9
`, 4.5555555556)
}
