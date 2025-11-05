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
	s := `2
100 30
40 10
`
	expect := 942477.796077000
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
1 1
9 7
1 4
10 7
`
	expect := 3983.539484752
	runSample(t, s, expect)
}
