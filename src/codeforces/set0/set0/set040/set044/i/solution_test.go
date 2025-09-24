package main

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, expect int) {
	ans := solve(n)
	if len(ans) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, ans)
	}

	seen := make(map[string]bool)

	for _, cur := range ans {
		slices.SortFunc(cur, func(a, b []int) int {
			return a[0] - b[0]
		})
		var flag int
		for _, x := range cur {
			sort.Ints(x)
			for _, i := range x {
				if (flag>>(i-1))&1 == 1 {
					t.Fatalf("Sample result %v, not correct", ans)
				}
				flag |= 1 << (i - 1)
			}
		}
		s := fmt.Sprintf("%v", cur)
		if seen[s] {
			t.Fatalf("Sample result %v, not correct", ans)
		}
		if flag != 1<<n-1 {
			t.Fatalf("Sample result %v, not correct", ans)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 15)
}
