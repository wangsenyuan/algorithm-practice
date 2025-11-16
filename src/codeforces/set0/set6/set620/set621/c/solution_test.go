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
	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 2
1 2
420 421
420420 420421
`
	expect := 4500.0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 5
1 4
2 3
11 14
`
	expect := 0.0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 3
3 3
2 4
1 1
`
	expect := 4666.666666666667

	// 1 * 3 * 1 = 3
	// 3 2 1 = 4000
	// 3 3 1 = 6000
	// 3 4 1 = 4000
	// 14000 / 3 = 

	runSample(t, s, expect)
}
