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
		t.Fatalf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 2 0.500000
1 2`
	runSample(t, s, 3.25)
}

func TestSample2(t *testing.T) {
	s := `4 3 0.4
4 3 1 2`
	runSample(t, s, 6.6312)
}
