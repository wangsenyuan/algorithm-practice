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
	if math.Abs(res-expect)/max(1, expect) > 1e-9 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
0 0
0 1
1 1
1 0
`
	runSample(t, s, 3.4142135624)
}

func TestSample2(t *testing.T) {
	s := `7
0 0
0 1000
1 1000
1000 999
1001 1
1000 0
1 -1
`
	runSample(t, s, 6825.7219029375)
}
