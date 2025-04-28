package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	n, likes, res := process(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	pos := make([]int, n)
	for i, cur := range res {
		for _, v := range cur {
			pos[v-1] = i
		}
	}

	for _, cur := range likes {
		u, v := cur[0], cur[1]
		u--
		v--
		if pos[u] != pos[v] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 0
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 4
1 2
2 3
3 4
5 6
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 3
1 2
2 3
1 3
`
	expect := true
	runSample(t, s, expect)
}
