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
	if math.Abs(res-expect) > 1e-6 && math.Abs(res-expect)/math.Max(1, math.Abs(expect)) > 1e-6 {
		t.Fatalf("Sample expect %.15f, but got %.15f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
1000 600 1200
`, 256.735020470879931)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
600 1000 1200
`, 261.423219407873376)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
100
`, -1100.000000000000000)
}
