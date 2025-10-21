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
	s := `3
4 7
8 10
5 5
`
	expect := 5.75
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
2 5
3 4
1 6
`
	expect := 3.5
	runSample(t, s, expect)
}
