package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, K, V, ok, res := drive(reader)
	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if !ok {
		return
	}
	n := len(a)
	if len(res) > n+5 {
		t.Fatalf("Sample result %v, is too long", res)
	}
	for _, cur := range res {
		cnt, x, y := cur[0], cur[1]-1, cur[2]-1
		if x == y || x < 0 || x >= n || y < 0 || y >= n {
			t.Fatalf("Sample result %v, is invalid", res)
		}
		w := a[x]
		if w/K >= cnt {
			a[y] += cnt * K
			a[x] -= cnt * K
		} else {
			// w / K < cnt
			a[y] += w
			a[x] = 0
		}
	}
	if slices.Contains(a, V) {
		return
	}
	t.Fatalf("Sample result %v, is invalid", res)
}

func TestSample1(t *testing.T) {
	s := `2 3 5
2 3`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 3 4
2 3
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 2 0
1 3 5 7 9
`
	expect := true
	runSample(t, s, expect)
}


func TestSample4(t *testing.T) {
	s := `8 4 20
3 3 3 3 3 3 3 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `8 6 32
3 6 4 4 4 4 4 4
`
	expect := true
	runSample(t, s, expect)
}

func TestK1AlreadyHasTarget(t *testing.T) {
	s := `2 1 4
2 4
`
	expect := true
	runSample(t, s, expect)
}

func TestNeedPartialKChunks(t *testing.T) {
	s := `2 2 11
5 9
`
	expect := true
	runSample(t, s, expect)
}

func TestZeroAlwaysPossible(t *testing.T) {
	s := `2 3 0
1 1
`
	expect := true
	runSample(t, s, expect)
}
