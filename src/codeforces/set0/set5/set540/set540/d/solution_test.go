package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []float64) {
	r := bufio.NewReader(strings.NewReader(s))
	res := drive(r)

	for i := range res {
		if math.Abs((res[i] - expect[i])) > 1e-9 {
			t.Fatalf("Sample %s, expect %v, but got %v", s, expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `2 2 2`
	expect := []float64{0.3333333333333333, 0.3333333333333333, 0.3333333333333333}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 1 2`
	expect := []float64{0.150000000000, 0.300000000000, 0.550000000000}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 1 3`
	expect := []float64{0.057142857143, 0.657142857143, 0.285714285714}
	runSample(t, s, expect)
}
