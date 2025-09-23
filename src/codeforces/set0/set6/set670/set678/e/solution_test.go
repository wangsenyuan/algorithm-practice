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
	if math.Abs(res-expect)/max(1.0, expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
0.0 0.5 0.8
0.5 0.0 0.4
0.2 0.6 0.0
`
	expect := 0.68
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1
0.0
`
	expect := 1.0
	runSample(t, s, expect)
}
