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
	if math.Abs(res-expect) > 1e-9 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1 0
1
`
	var expect float64
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 1
1 2
1 2 1
`
	expect := 3.0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 6
13 56 73 98 17
1 2 56
1 3 29
1 4 42
2 3 95
2 4 88
3 4 63
`
	expect := 2.965517241379311
	runSample(t, s, expect)
}
