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
		t.Fatalf("Sample expect %.10f, but got %.10f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 3 5
3 1 45
5 1 45
`
	expect := 2.0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 0 1
1 1 30
`
	expect := 0.732050808
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 0 1
1 1 45
`
	expect := 1.0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1 0 2
0 2 90
`
	expect := 2.0
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5 -10000 100000
993 94 2
-503 76 2
986 4 2
-312 21 1
338 6 2
`
	expect := 110000.000000000
	runSample(t, s, expect)
}
