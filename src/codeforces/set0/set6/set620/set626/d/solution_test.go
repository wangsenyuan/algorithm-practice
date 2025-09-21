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
	if math.Abs(res-expect)/max(1, expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
1 2
`
	expect := 0.0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 2 10
`
	expect := 0.0740740741
	runSample(t, s, expect)
}
