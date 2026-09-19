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
	if math.Abs(res-expect)/max(1.0, expect) > 1e-4 {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1 100
1
40
`, 40.0)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 100
1 1
25 30
`, 50.0)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 100
1 1
60 60
`, 100.0)
}
