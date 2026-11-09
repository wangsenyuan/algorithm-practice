package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	diff := math.Abs(res - expect)
	if diff > 1e-4*math.Max(1, math.Abs(expect)) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 2
100 3 5 7
`, 5.7428571429)
}

func TestSample2(t *testing.T) {
	runSample(t, `6 2
1 2 4 8 16 32
`, 8.5)
}
