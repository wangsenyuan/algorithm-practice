package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	y, x := drive(reader)
	mx := slices.Max(x)
	if mx != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, x)
	}

	n := len(x)

	slices.Sort(x)
	x = slices.Compact(x)
	if len(x) != n {
		t.Fatalf("Sample result %v, not by unique numbers ", x)
	}
	slices.Sort(y)
	marked := make([]bool, n)

	for range 30 {
		var l int
		for i := range n {
			for l < n && (marked[l] || x[l] < y[i]) {
				l++
			}
			if l < n && !marked[l] && y[i] == x[l] {
				marked[l] = true
				l++
			}
			y[i] /= 2
		}
	}

	for i := range n {
		if !marked[i] {
			t.Fatalf("Sample result %v, not correct, as %d-th number %d is not found", x, i, x[i])
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 2 3 4 5
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
15 14 3 13 1 12
`
	expect := 14
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
9 7 13 17 5 11
`
	expect := 6
	runSample(t, s, expect)
}
