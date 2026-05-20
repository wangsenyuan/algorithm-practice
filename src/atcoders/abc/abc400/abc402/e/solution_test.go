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
100 1 50
200 1 20
1000 1 1
`
	expect := 95.0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 32
500 9 57
300 4 8
300 3 32
300 7 99
100 8 69
`
	expect := 953.976967020096
	runSample(t, s, expect)
}
