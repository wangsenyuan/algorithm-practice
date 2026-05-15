package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	cubes, best, res := drive(reader)
	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
	var sum int
	for _, i := range res {
		sum += cubes[i-1][1]
	}
	if sum != best {
		t.Fatalf("Sample result %v, not getting the best %d", res, best)
	}
	first, second := -1, -1
	first = cubes[res[0]-1][0]
	if len(res) > 1 {
		second = cubes[res[1]-1][0]
	}
	if first == second {
		t.Fatalf("Sample result %v, first and second are the same %d", res, first)
	}
	for pos, i := range res {
		if pos%2 == 0 && first != cubes[i-1][0] {
			t.Fatalf("Sample result %v, first is not the color of the cube %d", res, first)
		}
		if pos%2 == 1 && second != cubes[i-1][0] {
			t.Fatalf("Sample result %v, second is not the color of the cube %d", res, second)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4
1 2
1 3
2 4
3 3
`
	expect := 9
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
1 1
2 1
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
2 1000000000
2 1000000000
2 1000000000
1 1
`
	expect := 2000000001
	runSample(t, s, expect)
}
