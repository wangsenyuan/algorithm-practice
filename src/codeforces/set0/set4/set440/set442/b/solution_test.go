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
	if math.Abs(res-expect) > 1e-7 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
0.1 0.2 0.3 0.8
`
	runSample(t, s, 0.8)
}

func TestSample2(t *testing.T) {
	runSample(t, "2\n0.1 0.2", 0.26)
}
