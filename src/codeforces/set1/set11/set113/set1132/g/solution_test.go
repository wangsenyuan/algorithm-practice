package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6 4
1 5 2 5 3 6
`, []int{2, 2, 3})
}

func TestSample2(t *testing.T) {
	runSample(t, `7 6
4 5 2 5 3 6 6
`, []int{3, 3})
}

func TestSingleElementWindows(t *testing.T) {
	runSample(t, `2 1
1 2
`, []int{1, 1})
}

func TestIncreasingWindows(t *testing.T) {
	runSample(t, `3 2
1 2 3
`, []int{2, 2})
}

func TestExhaustiveSmall(t *testing.T) {
	for n := 1; n <= 7; n++ {
		a := make([]int, n)
		var dfs func(int)
		dfs = func(i int) {
			if i == n {
				for k := 1; k <= n; k++ {
					expect := bruteForce(k, a)
					res := solve(k, slices.Clone(a))
					if !slices.Equal(res, expect) {
						t.Fatalf("n=%d k=%d a=%v, expect %v, but got %v", n, k, a, expect, res)
					}
				}
				return
			}
			for x := 1; x <= 4; x++ {
				a[i] = x
				dfs(i + 1)
			}
		}
		dfs(0)
	}
}

func bruteForce(k int, a []int) []int {
	n := len(a)
	ans := make([]int, n-k+1)
	for l := 0; l+k <= n; l++ {
		r := l + k
		for i := l; i < r; i++ {
			cur := i
			cnt := 1
			for {
				nxt := -1
				for j := cur + 1; j < r; j++ {
					if a[cur] < a[j] {
						nxt = j
						break
					}
				}
				if nxt < 0 {
					break
				}
				cur = nxt
				cnt++
			}
			ans[l] = max(ans[l], cnt)
		}
	}
	return ans
}
