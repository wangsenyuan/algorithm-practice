package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []float64) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	for i := range 2 {
		if math.Abs(res[i]-expect[i])/max(expect[i], 1.0) > 1e-7 {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `7 3
7 15 0 19 10 5 12
2
1 6
7 4
`
	expect := []float64{5, 15.5}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2
10 8 1 17
2
2 3
3 2
`
	expect := []float64{4.5, 13.5}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `11 3
60 100 84 74 19 77 36 48 70 18 63
4
3 7 11
5 9 2
2 9 5
8 10 1
`
	expect := []float64{42, 63}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 2
1 2 3
2
1
2
`
	expect := []float64{1, 2}
	runSample(t, s, expect)
}
