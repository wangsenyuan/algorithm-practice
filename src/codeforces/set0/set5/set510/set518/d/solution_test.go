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
	runSample(t, "1 0.50 1", 0.5)
}

func TestSample2(t *testing.T) {
	runSample(t, "1 0.50 4", 0.9375)
}

func TestSample3(t *testing.T) {
	runSample(t, "4 0.20 2", 0.4)
}
