package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
1 1
`, []int{1})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
2 1
`, []int{3})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
2 2
`, []int{0})
}

func TestSample4(t *testing.T) {
	runSample(t, `1
3 1
`, []int{6})
}

func TestSample5(t *testing.T) {
	runSample(t, `1
3 2
`, []int{2})
}

func TestSample6(t *testing.T) {
	runSample(t, `1
3 3
`, []int{0})
}

func TestSolveUsesPrecomputedAnswers(t *testing.T) {
	allocs := testing.AllocsPerRun(5, func() {
		_ = solve(100, 5)
	})
	if allocs != 0 {
		t.Fatalf("solve must be an O(1), allocation-free lookup, got %.0f allocations", allocs)
	}
}

func TestSmallAgainstBruteForce(t *testing.T) {
	for n := 1; n <= 8; n++ {
		for k := 1; k <= n; k++ {
			type state struct {
				start int
				left  int
				used  int
			}
			memo := make(map[state]int)
			var brute func(int, int, int) int
			brute = func(start, left, used int) int {
				if left == 0 {
					return 1
				}
				key := state{start, left, used}
				if v, ok := memo[key]; ok {
					return v
				}
				res := 0
				for a := start; a <= n; a++ {
					for b := a; b <= n; b++ {
						d := b - a
						if used>>d&1 == 0 {
							res += brute(b+1, left-1, used|1<<d)
						}
					}
				}
				memo[key] = res
				return res
			}

			expect := brute(1, k, 0) % mod
			if got := solve(n, k); got != expect {
				t.Fatalf("n=%d k=%d: got %d, expect %d", n, k, got, expect)
			}
		}
	}
}
