package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	f, m, g, h := drive(reader)
	if m > 0 != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, m > 0)
	}
	if !expect {
		return
	}
	n := len(f)
	for i := 1; i <= m; i++ {
		j := h[i-1]
		x := g[j-1]
		if x != i {
			// g(h(i)) = i
			t.Fatalf("Sample result %v, %v, not correct", g, h)
		}
	}

	for i := 1; i <= n; i++ {
		j := g[i-1]
		x := h[j-1]
		if x != f[i-1] {
			// h(g(i)) = f(i)
			t.Fatalf("Sample result %v, %v, not correct", g, h)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 2 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
2 2 2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
2 1
`
	expect := false
	runSample(t, s, expect)
}