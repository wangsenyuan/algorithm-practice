package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	l, a := drive(reader)
	if len(a) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, len(a) > 0)
	}
	if !expect {
		return
	}
	n := len(a)
	for i := 0; i+1 < len(l); i++ {
		v := l[i] - 1
		w := l[i+1] - 1
		if (v+a[v])%n != w {
			t.Fatalf("Sample result %v, not correct", a)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4 5
2 3 1 4 4
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
3 1 2
`
	expect := false
	runSample(t, s, expect)
}
