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
	if len(res) != 1 || math.Abs(res[0]-expect) > 1e-6 {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
0 0 -2 2
-1 -1 4 4
`, 1.0)
}

func TestSample2(t *testing.T) {
	runSample(t, `1
4 0 2 0
6 0 8 0
`, 2.0)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
1 0 1 1
-1 0 1 1
`, 0.0)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
-8 9 2 6
-10 -10 17 20
`, 1.783905950993199)
}
