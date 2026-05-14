package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	p, x, ok, res := drive(reader)
	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if !ok {
		return
	}

	y := make([]byte, len(x))
	for i := range len(x) {
		y[i] = '0'
	}
	for _, e := range res {
		l, r := e[0]-1, e[1]-1
		lo := min(p[l], p[r])
		hi := max(p[l], p[r])
		for i := l + 1; i < r; i++ {
			if p[i] > lo && p[i] < hi {
				y[i] = '1'
			}
		}
	}

	for i := range len(x) {
		if x[i] == '1' && y[i] == '0' {
			t.Fatalf("Sample result %v, is incorrect, it leads to %s, but expect %s", res, y, x)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 2 3
010
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
3 4 2 1 5
11111
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
1 3 2 4 6 5
001100
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6
6 2 3 4 5 1
110110
`
	expect := false
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5
2 1 4 3 5
00000
`
	expect := true
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `5
2 5 3 1 4
00100
`
	expect := true
	runSample(t, s, expect)
}



