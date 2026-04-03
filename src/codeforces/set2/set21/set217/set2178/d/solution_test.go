package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	h, m, res := drive(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, len(res) > 0)
	}
	if !expect {
		return
	}
	a := slices.Clone(h)
	n := len(h)
	marked := make([]bool, n)
	for _, cur := range res {
		x, y := cur[0]-1, cur[1]-1
		if marked[x] {
			t.Fatalf("Sample result %v, not valid, as %d already attacked someone", res, x+1)
		}
		if h[x] <= 0 || h[y] <= 0 {
			t.Fatalf("Sample result %v, not valid, as %d or %d is not alive", res, x+1, y+1)
		}
		marked[x] = true
		h[x] -= a[y]
		h[y] -= a[x]
	}
	var cnt int
	for _, v := range h {
		if v > 0 {
			cnt++
		}
	}
	if cnt != m {
		t.Fatalf("Sample result %v, not valid, as %d elves are alive (want %d)", res, cnt, m)
	}
	if cnt >= 2 {
		for i := range n {
			if h[i] > 0 && !marked[i] {
				t.Fatalf("Sample result %v, not valid: elf %d is alive but never attacked", res, i+1)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 2
1 4 2 3`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 2
6 7`, false)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 0
1 2 3`, true)
}

func TestSample4(t *testing.T) {
	runSample(t, `3 1
1 2 3`, true)
}

func TestSample5(t *testing.T) {
	runSample(t, `3 2
1 2 3`, false)
}

func TestSample6(t *testing.T) {
	runSample(t, `4 1
2 3 4 5`, true)
}

func TestSample7(t *testing.T) {
	runSample(t, `6 0
998244353 1000000000 314159265 676767677 999999999 987654321`, true)
}
