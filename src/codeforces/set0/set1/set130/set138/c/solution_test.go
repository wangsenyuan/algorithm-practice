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
	if math.IsNaN(res) || math.Abs(res-expect)/max(1, expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1 1
2 2 50 50
1 1
`
	expect := 0.5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 1
2 2 50 50
4 2 50 50
3 1
`
	expect := 0.25
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 2
-8 4 66 9
-2 3 55 43
3 8
7 9
`
	expect := 17.0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 3
-5 6 0 100
5 6 100 0
6 3
0 4
-6 3
`
	expect := 6.0
	runSample(t, s, expect)
}
