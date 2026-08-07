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
	runSample(t, `1
1
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
1 3 5 2 4
`, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, `6
1 2 4 5 3 6
`, 3)
}

func TestSample4(t *testing.T) {
	runSample(t, `7
1 3 5 7 2 4 6
`, 0)
}

func TestRegressionPrefixViolationCanDisappear(t *testing.T) {
	runSample(t, `5
1 2 4 3 5
`, 4)
}

func TestSmallPermutations(t *testing.T) {
	for n := 1; n <= 7; n++ {
		p := make([]int, n)
		for i := range n {
			p[i] = i + 1
		}
		for {
			expect := bruteForce(p)
			res := solve(append([]int(nil), p...))
			if res != expect {
				t.Fatalf("n=%d p=%v expect %d, but got %d", n, p, expect, res)
			}
			if !nextPermutation(p) {
				break
			}
		}
	}
}

func bruteForce(p []int) int {
	n := len(p)
	var res int
	for s := range n {
		seen := make([]bool, n+2)
		blocks := 0
		ok := true
		for i := range n {
			v := p[(s+i)%n]
			if !seen[v-1] {
				blocks++
			}
			if seen[v+1] {
				blocks--
			}
			seen[v] = true
			if blocks > 2 {
				ok = false
				break
			}
		}
		if ok {
			res++
		}
	}
	return res
}

func nextPermutation(a []int) bool {
	i := len(a) - 2
	for i >= 0 && a[i] >= a[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := len(a) - 1
	for a[j] <= a[i] {
		j--
	}
	a[i], a[j] = a[j], a[i]
	for l, r := i+1, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
	return true
}
