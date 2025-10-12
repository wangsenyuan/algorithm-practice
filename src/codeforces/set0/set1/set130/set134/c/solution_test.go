package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, res := drive(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, len(res) > 0)
	}
	if !expect {
		return
	}
	// n := len(a)
	// 怎么验证这个问题呢？ 貌似有点难呐～
	n := len(a)
	before := slices.Clone(a)
	after := make([]int, n)
	for _, cur := range res {
		u, v := cur[0]-1, cur[1]-1
		if before[u] == 0 || before[v] == 0 {
			t.Fatalf("Sample result %v, swap %v, not possible", res, cur)
		}
		before[u]--
		before[v]--
		after[u]++
		after[v]++
	}

	if !slices.Equal(a, after) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 8
2 2 2 2`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 12
1 1 2 2 3 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 5
0 0 0 0 5
`
	expect := false
	runSample(t, s, expect)
}
