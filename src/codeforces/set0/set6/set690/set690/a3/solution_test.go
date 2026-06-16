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
2 1
1
2 1
2
2 2
1
2 2
2
`
	runSample(t, s, []int{1, 2, 2, 1})
}

func TestStrategyAlwaysHasCorrectGuess(t *testing.T) {
	for n := 2; n <= 6; n++ {
		nums := make([]int, n)
		var dfs func(int)
		dfs = func(pos int) {
			if pos == n {
				correct := 0
				for rank := 1; rank <= n; rank++ {
					seen := make([]int, 0, n-1)
					for i, x := range nums {
						if i+1 != rank {
							seen = append(seen, x)
						}
					}
					if solve(n, rank, seen) == nums[rank-1] {
						correct++
					}
				}
				if correct == 0 {
					t.Fatalf("n=%d assignment=%v has no correct guess", n, nums)
				}
				return
			}
			for x := 1; x <= n; x++ {
				nums[pos] = x
				dfs(pos + 1)
			}
		}
		dfs(0)
	}
}
