package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	requests, tables, best, accepted := drive(bufio.NewReader(strings.NewReader(s)))

	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}

	var sum int

	for _, cur := range accepted {
		x, y := cur[0]-1, cur[1]-1
		if requests[x][0] > tables[y] {
			t.Fatalf("Sample result %v not correct at %v", accepted, cur)
		}
		sum += requests[x][1]
	}

	if sum != best {
		t.Fatalf("Sample result %v, profit %d, instead of %d", accepted, sum, best)
	}

}

func TestSample1(t *testing.T) {
	s := `3
10 50
2 100
5 30
3
4 6 9
`
	expect := 130
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
10 10
3 5
5 8
3
3 4 10
`
	expect := 15
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
10 10
3 5
5 8
3
3 4 10
`
	expect := 15
	runSample(t, s, expect)
}
