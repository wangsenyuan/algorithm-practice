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
	if math.Abs(res-expect) > 1e-5 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 3
10 20 30 40
1 3
2 3
4 3
`
	expect := 16.6666667
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
10 20 30
1 2
2 3
3 1
`
	expect := 13.333333
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7 8
40 20 10 30 20 50 40
1 2
2 3
3 4
4 5
5 6
6 7
1 4
5 7
`
	expect := 18.571429
	runSample(t, s, expect)
}
