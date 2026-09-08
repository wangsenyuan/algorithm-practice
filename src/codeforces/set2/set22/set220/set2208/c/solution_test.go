package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	t.Helper()
	t.Skip("solve TODO")
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if math.Abs(res-expect)/max(1, math.Abs(expect)) > 1e-6 {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2
10 0
20 5
`, 30)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
10 5
10 80
20 5
`, 29)
}
