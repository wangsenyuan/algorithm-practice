package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	p, ok, res := drive(reader)
	if ok != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	n := len(p)
	if len(res) != n-1 {
		t.Fatalf("Sample result %v, not correct", res)
	}
	for i := range n {
		p[i]--
	}
	type pair struct {
		first  int
		second int
	}
	mem := make(map[pair]bool)
	for _, cur := range res {
		u, v := cur[0]-1, cur[1]-1
		u, v = min(u, v), max(u, v)
		mem[pair{u, v}] = true
	}

	for _, cur := range res {
		u, v := cur[0]-1, cur[1]-1
		u, v = min(u, v), max(u, v)
		if !mem[pair{u, v}] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4
4 3 2 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
3 1 2
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
3 4 1 2
`
	expect := true
	runSample(t, s, expect)
}
