package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	points, best, rect := drive(reader)
	if best != expect {
		t.Fatalf("Sample expect %d, but got %d, %v", expect, best, rect)
	}

	var sum int
	for _, p := range points {
		x, y := p[0], p[1]
		x, y = min(x, y), max(x, y)
		if x >= rect[0] && y <= rect[2] {
			sum += p[2]
		}
	}

	sum -= rect[2] - rect[0]
	if sum != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, sum)
	}
}

func TestSample1(t *testing.T) {
	s := `6
0 0 2
1 0 -5
1 1 3
2 3 4
1 4 -4
3 1 -1
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
3 3 0
3 3 -3
0 2 -1
3 1 3
0 0 -2
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
0 0 -1
0 1 1
1 0 -2
1 1 -1
`
	expect := 0
	runSample(t, s, expect)
}

