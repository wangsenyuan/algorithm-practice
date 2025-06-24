package main

import (
	"bufio"
	"sort"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, m, k, sum, res := process(reader)
	if len(res) != k-1 {
		t.Fatalf("Sample result %v, not correct", res)
	}
	expect := readNum(reader)

	if expect != sum {
		t.Fatalf("Sample expect %d, but got %d", expect, sum)
	}
	res = append(res, len(a))
	prev := 0
	var cur int
	for _, j := range res {
		tmp := a[prev:j]
		if len(tmp) < m {
			t.Fatalf("Sample result split wrong size sub-array %v", tmp)
		}
		sort.Ints(tmp)
		for i := len(tmp) - m; i < len(tmp); i++ {
			cur += tmp[i]
		}
		prev = j
	}
	if cur != sum {
		t.Fatalf("Sample result expect  %d, but got %d", expect, cur)
	}
}

func TestSample1(t *testing.T) {
	s := `9 2 3
5 2 5 2 4 1 1 3 2
21`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `6 1 4
4 1 3 2 2 3
12`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `2 1 2
-1000000000 1000000000
0`
	runSample(t, s)
}
