package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, edges, res := drive(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	a := res[0]
	b := res[1]

	get := func(arr []int, u int, v int) int {
		if arr[u] > arr[v] {
			return 1
		}
		if arr[u] < arr[v] {
			return -1
		}
		return 0
	}

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		x := get(a, u, v)
		y := get(b, u, v)
		if x != y {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	slices.Sort(a)
	a = slices.Compact(a)
	if len(a) != n {
		t.Fatalf("Sample result %v, not correct, it should have %d distinct elements for a", res, n)
	}

	slices.Sort(b)
	b = slices.Compact(b)
	if len(b) != n-1 {
		t.Fatalf("Sample result %v, not correct, it should have %d distinct elements for b", res, n-1)
	}
}

func TestSample1(t *testing.T) {
	s := `1 0`
	expect := false
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1
1 2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 3
1 2
1 3
2 4
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 0`
	expect := true
	runSample(t, s, expect)
}
