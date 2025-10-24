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
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 4
1 1 1 1
`
	expect := 3.0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 18
4 4 4 2 2 2
`
	expect := 18.0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 5
2 3
`
	expect := 4.5
	runSample(t, s, expect)
}
