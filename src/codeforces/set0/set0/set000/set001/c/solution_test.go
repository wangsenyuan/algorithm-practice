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
	if math.Abs(res-expect) > 1e-7 {
		t.Fatalf("Sample expect %.8f, but got %.8f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `0.000000 0.000000
1.000000 1.000000
0.000000 1.000000
`, 1)
}

func TestRegularPentagonVertices(t *testing.T) {
	runSample(t, `1.000000 0.000000
-0.809016994375 0.587785252292
0.309016994375 -0.951056516295
`, 2.377641290737884)
}

func TestSample2(t *testing.T) {
	runSample(t, `77.145533 85.041789
67.452820 52.513188
80.503843 85.000149
`, 1034.7083664592612)
}
