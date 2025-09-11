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
	s := `4
1 2
1 3
2 4`
	expect := 1.5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 2
1 3
3 4
2 5
`
	expect := 2.0
	runSample(t, s, expect)
}
