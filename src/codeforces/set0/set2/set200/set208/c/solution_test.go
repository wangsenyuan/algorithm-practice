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
	if math.Abs(res-expect) > 1e-7 {
		t.Fatalf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 4
1 2
2 4
1 3
3 4
`
	expect := 1.0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `11 14
1 2
1 3
2 4
3 4
4 5
4 6
5 11
6 11
1 8
8 9
9 7
11 7
1 10
10 4
`
	expect := 1.714285714286

	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `13 22
10 13
9 13
13 4
2 13
6 13
5 1
1 8
9 1
13 5
1 7
6 1
1 12
4 1
13 8
1 3
10 1
13 12
1 11
13 11
1 2
13 3
13 7
`
	expect := 1.0
	runSample(t, s, expect)
}