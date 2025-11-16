package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, b, best, res := drive(reader)

	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}

	for _, cur := range res {
		i, j := cur[0]-1, cur[1]-1
		a[i], b[j] = b[j], a[i]
	}

	var s1, s2 int
	for _, v := range a {
		s1 += v
	}
	for _, v := range b {
		s2 += v
	}

	if abs(s1-s2) != best {
		t.Fatalf("Sample result %v, got different best %d, s1 = %d, s2 = %d", res, best, s1, s2)
	}
}

func TestSample1(t *testing.T) {
	s := `5
5 4 3 2 1
4
1 1 1 1
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 2 3 4 5
1
15
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
1 2 3 4 5
4
1 2 3 4
`
	expect := 1
	runSample(t, s, expect)
}
