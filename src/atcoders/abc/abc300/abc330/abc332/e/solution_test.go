package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	res := process(bufio.NewReader(strings.NewReader(s)))
	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 3
3 5 3 6 3
`
	expect := 0.888888888888889
	runSample(t, s, expect)
}
