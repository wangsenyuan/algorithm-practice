package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	c, res := drive(reader)
	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, len(res))
	}
	n := len(c)
	a := make([]int, n)
	pos := make([]int, n+1)
	for i := range n {
		a[i] = i + 1
		pos[a[i]] = i
	}

	for _, cur := range res {
		i, j := cur[0], cur[1]
		if pos[i]-1 != pos[j] {
			t.Fatalf("Sample result %v, not valid, %d can't overtake %d", res, i, j)
		}

		pos[i], pos[j] = pos[j], pos[i]

		a[pos[i]] = i
		a[pos[j]] = j
	}

	if !slices.Equal(a, c) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
2 3 1
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
1 2
`
	expect := 2
	runSample(t, s, expect)
}
