package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if math.Abs(res-expect) > 1e-6 {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1 7 6 7
4
4 2 4 12 3 12 3 2
`, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, `-1 0 2 0
4
0 0 1 0 1 1 0 1
`, 3)
}

func TestMissesIsland(t *testing.T) {
	runSample(t, `0 10 5 10
4
0 0 1 0 1 1 0 1
`, 5)
}

func TestAroundIsCheaper(t *testing.T) {
	runSample(t, `-1 -1 3 3
4
0 0 2 0 2 2 0 2
`, 4+2*math.Sqrt(2))
}
