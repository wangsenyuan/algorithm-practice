package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := drive(reader)
	if math.Abs(ans-expect)/max(1, expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `2 3
-1 3
0 0
-1 -4`
	expect := -2.5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 2
0 2
5 0
`
	expect := 7.0
	runSample(t, s, expect)
}
