package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, res := drive(reader)

	n := len(a)

	check := func(res []int) []int {
		cnt := make([]int, 4)
		d := make([]int, 4)
		c := make([]int, 4)
		for i := range 4 {
			d[i] = inf
			c[i] = -inf
		}

		for i := range n {
			x := res[i]
			if x > 0 {
				cnt[x-1]++
				d[x-1] = min(d[x-1], a[i])
				c[x-1] = max(c[x-1], a[i])
			} else {
				d[3] = min(d[3], a[i])
				c[3] = max(c[3], a[i])
			}
		}

		w := slices.Max(cnt[:3])
		v := slices.Min(cnt[:3])
		if v*2 < w || v == 0 {
			t.Fatalf("Sample result %v, not valid", res)
		}
		p := make([]int, 3)
		p[0] = d[0] - c[1]
		p[1] = d[1] - c[2]
		p[2] = d[2] - max(c[3], 0)
		return p
	}

	dx := check(expect)
	dy := check(res)

	if !slices.Equal(dx, dy) {
		t.Fatalf("Sample expect %v(%v), but got %v(%v)", expect, dx, res, dy)
	}
}

func TestSample1(t *testing.T) {
	s := `4
1 2 3 4
`
	expect := []int{3, 3, 2, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
1 4 3 1 1 2
`
	expect := []int{-1, 1, 2, -1, -1, 3}
	runSample(t, s, expect)
}
