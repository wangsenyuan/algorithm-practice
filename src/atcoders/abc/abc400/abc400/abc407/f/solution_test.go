package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
5 3 4 2
`
	// 5 贡献 1,2,3,4
	// 3 贡献 1
	// 4 贡献 1,2(2), 3
	// 2 贡献 1
	runSample(t, s, []int{14, 13, 9, 5})
}

func TestSample2(t *testing.T) {
	s := `8
2 0 2 5 0 5 2 4
`
	runSample(t, s, []int{20, 28, 27, 25, 20, 15, 10, 5})
}

func TestSample3(t *testing.T) {
	s := `11
9203973 9141294 9444773 9292472 5507634 9599162 497764 430010 4152216 3574307 430010
`
	runSample(t, s, []int{
		61273615, 68960818, 69588453, 65590626, 61592799, 57594972,
		47995810, 38396648, 28797486, 19198324, 9599162,
	})
}

func TestSinglePeakCanCoverWholeArray(t *testing.T) {
	s := `3
1 3 2
`
	runSample(t, s, []int{6, 6, 3})
}

func TestAgainstBruteForceSmall(t *testing.T) {
	for n := 1; n <= 6; n++ {
		a := make([]int, n)
		var dfs func(int)
		dfs = func(pos int) {
			if pos == n {
				expect := bruteForce(a)
				res := solve(a)
				if !reflect.DeepEqual(res, expect) {
					t.Fatalf("solve(%v) = %v, expect %v", a, res, expect)
				}
				return
			}
			for v := 0; v <= 3; v++ {
				a[pos] = v
				dfs(pos + 1)
			}
		}
		dfs(0)
	}
}

func bruteForce(a []int) []int {
	n := len(a)
	res := make([]int, n)
	for k := 1; k <= n; k++ {
		for l := 0; l+k <= n; l++ {
			best := a[l]
			for r := l; r < l+k; r++ {
				best = max(best, a[r])
			}
			res[k-1] += best
		}
	}
	return res
}
