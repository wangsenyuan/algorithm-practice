package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if math.Abs(res-expect) > eps {
		t.Fatalf("Sample expect %.6f, but got %.6f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 4
0 5 6 5
3
0 0 0 4
1 1 4 1
6 0 6 4
`
	expect := 19.0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 10
0 0 10 10
1
5 0 5 9
`
	var expect float64 = -1
	runSample(t, s, expect)
}
