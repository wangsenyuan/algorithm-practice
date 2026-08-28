package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2
1
`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
2 2
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `3
1 1
`, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, `4
2 3 2
`, 0)
}

func TestSample5(t *testing.T) {
	runSample(t, `5
3 3 4 2
`, 2)
}

func TestSample6(t *testing.T) {
	runSample(t, `2
2
`, 0)
}

func TestSample7(t *testing.T) {
	runSample(t, `3
1 2
`, 2)
}

func TestSample8(t *testing.T) {
	runSample(t, `4
3 3 3
`, 4)
}

func TestSample9(t *testing.T) {
	runSample(t, `5
4 4 4 4
`, 12)
}

func TestSample10(t *testing.T) {
	runSample(t, `4
2 1 2
`, 0)
}

func TestSample11(t *testing.T) {
	runSample(t, `6
3 3 5 5 5
`, 8)
}

func TestSolveMatchesAllPermutationsUpToSix(t *testing.T) {
	for n := 2; n <= 6; n++ {
		want := make(map[int]int)
		p := make([]int, n)
		used := make([]bool, n+1)
		var buildPerm func(int)
		buildPerm = func(at int) {
			if at == n {
				code := 0
				for cut := 1; cut < n; cut++ {
					left, right := 0, 0
					for i, value := range p {
						if i < cut && value > left {
							left = value
						}
						if i >= cut && value > right {
							right = value
						}
					}
					if left > right {
						left = right
					}
					code = code*n + left - 1
				}
				want[code]++
				return
			}
			for value := 1; value <= n; value++ {
				if !used[value] {
					used[value] = true
					p[at] = value
					buildPerm(at + 1)
					used[value] = false
				}
			}
		}
		buildPerm(0)

		a := make([]int, n-1)
		var checkAll func(int)
		checkAll = func(at int) {
			if at == len(a) {
				code := 0
				for _, value := range a {
					code = code*n + value - 1
				}
				if got := solve(a); got != want[code] {
					t.Fatalf("n=%d, a=%v: expect %d, got %d", n, a, want[code], got)
				}
				return
			}
			for value := 1; value <= n; value++ {
				a[at] = value
				checkAll(at + 1)
			}
		}
		checkAll(0)
	}
}
