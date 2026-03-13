package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	x, y, p, best, res := drive(reader)

	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}

	f := func(i int, j int) int {
		return (x[i] + y[j]) % p
	}

	sum := f(0, 0)
	var a, b int
	for _, c := range res {
		if c == 'C' {
			a++
		} else {
			b++
		}
		sum += f(a, b)
	}

	if sum != best {
		t.Fatalf("Sample result %s, not getting the best %d", res, best)
	}
}

func TestSample1(t *testing.T) {
	s := `2 2 10
0 0
0 1
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3 10
0 2 0
0 0 2
`
	expect := 10
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 3 2
0 1 1
1 1 0
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 4 3
2 0 0 0
0 0 0 2
`
	expect := 13
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5 3 2
0 1 0 0 0
0 1 0
`
	expect := 4
	runSample(t, s, expect)
}
