package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 2
0.500000 0.250000 0.250000
	`, 0.62500000)
}

func TestSample2(t *testing.T) {
	runSample(t, `9 9
0.100000 0.100000 0.100000 0.100000 0.100000 0.100000 0.100000 0.100000 0.100000 0.100000
	`, 0.93687014)
}
