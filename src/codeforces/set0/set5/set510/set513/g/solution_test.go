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
	s := `3 1
1 2 3
`
	expect := 0.833333333333333
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 4
1 3 2
`
	expect := 1.458333333333334
	runSample(t, s, expect)
}
