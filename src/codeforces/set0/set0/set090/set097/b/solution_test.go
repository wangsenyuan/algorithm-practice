package main

import (
	"bufio"
	"cmp"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	slices.SortFunc(res, func(a, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	check := func(l int, r int) bool {
		if res[l][0] == res[r][0] || res[l][1] == res[r][1] {
			return true
		}
		// a[0] < b[0]

		for i := range len(res) {
			if i == l || i == r {
				continue
			}
			if res[i][0] >= res[l][0] && res[i][0] <= res[r][0] {
				if res[i][1] >= min(res[l][1], res[r][1]) && res[i][1] <= max(res[l][1], res[r][1]) {
					return true
				}
			}
		}
		return false
	}

	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if !check(i, j) {
				t.Fatalf("Sample result %v, causing conflict", res)
			}
		}
	}

	for i := 0; i+1 < len(res); i++ {
		if res[i][0] == res[i+1][0] && res[i][1] == res[i+1][1] {
			t.Fatalf("Sample result %v, has duplicates %v", res, res[i])
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4
-10 1
-2 -5
-1 -2
4 -3`
	runSample(t, s)
}
