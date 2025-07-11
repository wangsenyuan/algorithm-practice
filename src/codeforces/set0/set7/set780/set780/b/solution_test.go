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
	runSample(t, `3
7 1 3
1 2 1
`, 2.0)
}

func TestSample2(t *testing.T) {
	runSample(t, `4
5 10 3 2
2 3 2 4
`, 1.4)
}
