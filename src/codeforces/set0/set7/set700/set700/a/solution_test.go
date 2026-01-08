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
	s := `5 10 1 2 5`
	expect := 5.0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 6 1 2 1`
	expect := 4.7142857143
	runSample(t, s, expect)
}
