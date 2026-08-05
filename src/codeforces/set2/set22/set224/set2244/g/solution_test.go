package main

import (
	"bufio"
	"math/rand"
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
	runSample(t, `5
3 1 1 1 3
`, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
2 1 2
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `7
4 1 5 1 1 4 1
`, 8)
}

func bruteForce(a []int) int {
	n := len(a)
	best := 0
	for mask := 0; mask < 1<<n; mask++ {
		sum := 0
		ok := true
		for i := 0; i < n && ok; i++ {
			if mask>>i&1 == 0 {
				continue
			}
			sum += a[i]
			for j := 0; j < i; j++ {
				if mask>>j&1 == 1 && i-j <= max(a[i], a[j]) {
					ok = false
					break
				}
			}
		}
		if ok {
			best = max(best, sum)
		}
	}
	return best
}

func TestSolveAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for n := 1; n <= 12; n++ {
		for it := 0; it < 1000; it++ {
			a := make([]int, n)
			for i := range a {
				a[i] = rng.Intn(n + 2)
			}
			expect := bruteForce(a)
			if res := solve(a); res != expect {
				t.Fatalf("solve(%v) = %d, want %d", a, res, expect)
			}
		}
	}
}
