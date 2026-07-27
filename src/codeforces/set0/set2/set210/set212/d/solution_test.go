package main

import (
	"bufio"
	"math"
	"math/rand"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []float64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if len(res) != len(expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	for i := range expect {
		if math.Abs(res[i]-expect[i]) > 1e-9 {
			t.Errorf("Sample expect %v, but got %v", expect, res)
			return
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
3 2 1
3
1 2 3
`, []float64{2.0, 1.5, 1.0})
}

func TestSample2(t *testing.T) {
	runSample(t, `2
1 1
3
1 2 1
`, []float64{1.0, 1.0, 1.0})
}

func solveBruteForce(a []int, ks []int) []float64 {
	res := make([]float64, len(ks))
	for id, k := range ks {
		var sum int64
		for l := 0; l+k <= len(a); l++ {
			lowest := a[l]
			for i := l + 1; i < l+k; i++ {
				lowest = min(lowest, a[i])
			}
			sum += int64(lowest)
		}
		res[id] = float64(sum) / float64(len(a)-k+1)
	}
	return res
}

func TestSolveAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(212))
	for n := 1; n <= 30; n++ {
		for test := 0; test < 100; test++ {
			a := make([]int, n)
			for i := range a {
				a[i] = rng.Intn(6) + 1
			}

			ks := rng.Perm(n)
			for i := range ks {
				ks[i]++
			}
			expect := solveBruteForce(a, ks)
			result := solve(a, ks)
			for i := range expect {
				if math.Abs(result[i]-expect[i]) > 1e-9 {
					t.Fatalf("a = %v, k = %d, expect %.15f, but got %.15f",
						a, ks[i], expect[i], result[i])
				}
			}
		}
	}
}
