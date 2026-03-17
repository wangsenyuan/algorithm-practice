package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, a, ok, ans := drive(reader)
	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if !ok {
		return
	}
	if len(ans) != n {
		t.Fatalf("Sample expect %d, but got %d", n, len(ans))
	}
	row := make([]bool, n)
	col := make([]bool, n)
	for _, cur := range ans {
		r, c := cur[0]-1, cur[1]-1
		if a[r][c] == 'E' {
			t.Fatalf("Sample result %v, not correct", ans)
		}
		row[r] = true
		col[c] = true
	}
	for i := range n {
		for j := range n {
			if !row[i] && !col[j] {
				t.Fatalf("Sample result %v, not correct", ans)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3
.E.
E.E
.E.
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
EEE
E..
E.E
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
EE.EE
E.EE.
E...E
.EE.E
EE.EE
`
	expect := true
	runSample(t, s, expect)
}
