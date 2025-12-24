package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	p, x, res, a, assign := drive(reader)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}

	marked := make([]bool, len(x))
	var c, u int
	for i := range len(p) {
		if assign[i] == 0 {
			continue
		}
		j := assign[i] - 1
		if marked[j] {
			t.Fatalf("Sample result %v, not correct, it reuse %d-th socket", assign, j)
		}
		marked[j] = true
		for x[j] > p[i] {
			x[j] = (x[j] + 1) / 2
			u++
		}
		if x[j] < p[i] {
			t.Fatalf("Sample result %v, not correct, it can not connect %d-th computer to %d-th socket", assign, i, j)
		}
		c++
	}

	if c != expect[0] || u != expect[1] {
		t.Fatalf("Sample result %v, not correct, expect %d %d, but got %d %d", assign, expect[0], expect[1], c, u)
	}

	var w int
	for _, v := range a {
		w += v
	}

	if w != u {
		t.Fatalf("Sample result %v, not correct, expect %d, but got %d", assign, u, w)
	}
}

func TestSample1(t *testing.T) {
	s := `2 1
2 100
99
`
	expect := []int{1, 6}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2
1 1
2 2
`
	expect := []int{2, 2}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 1
2 100
99
`
	expect := []int{1, 6}
	runSample(t, s, expect)
}