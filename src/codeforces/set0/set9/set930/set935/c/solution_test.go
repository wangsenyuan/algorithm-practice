package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [3]float64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	for i := range 3 {
		if math.Abs(res[i]-expect[i])/math.Max(1, math.Abs(expect[i])) > 1e-6 {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 3 3 1 1
`, [3]float64{3.7677669529663684, 3.7677669529663684, 3.914213562373095})
}

func TestSample2(t *testing.T) {
	runSample(t, `10 5 5 5 15
`, [3]float64{5.0, 5.0, 10.0})
}

func TestLaptopAtCenter(t *testing.T) {
	runSample(t, `8 2 3 2 3
`, [3]float64{6.0, 3.0, 4.0})
}
