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
12
11 8
7 5
`
	expect := 10.0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1
1 4 1
2 5 3
`
	expect := -1.0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
2
4 6 3 3 5 6
2 6 3 6 5 3
`
	expect := 85.48
	runSample(t, s, expect)
}
