package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if math.Abs(res-expect) > 1e-9 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2
0 40 100 100
60 0 40 40
0 60 0 45
0 60 55 0
`, 1.75)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
0 0 100 0 100 0 0 0
100 0 100 0 0 0 100 100
0 0 0 100 100 0 0 0
100 100 0 0 0 0 100 100
0 100 0 100 0 0 100 0
100 100 100 100 100 0 0 0
100 0 100 0 0 100 0 0
100 0 100 0 100 100 100 0
`, 12.0)
}

func TestSample3(t *testing.T) {
	runSample(t, `2
0 21 41 26
79 0 97 33
59 3 0 91
74 67 9 0
`, 3.141592)
}
