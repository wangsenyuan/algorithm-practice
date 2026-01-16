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
	diff := math.Abs(res - expect)
	if diff/max(1.0, expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f (diff: %f)", expect, res, diff)
	}
}

func TestSample1(t *testing.T) {
	s := `2
40 10 50
60 20 30
`
	expect := 70.0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
50 30 80
35 25 70
40 10 90
`
	expect := 55.0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
1 1 2
2 2 3
3 3 4
`
	expect := 6.0
	runSample(t, s, expect)
}
