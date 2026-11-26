package main

import (
	"bufio"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	got := drive(reader)
	if !reflect.DeepEqual(got, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, got)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 5
1 2 3 4 5
2 3 5 4
3 3 5
1 2 5
2 1 3 3
1 1 3
`, []int64{8, 5})
}

func TestSample2(t *testing.T) {
	runSample(t, `10 10
6 9 6 7 6 1 10 10 9 5
1 3 9
2 7 10 9
2 5 10 8
1 4 7
3 3 7
2 7 9 9
1 2 4
1 6 6
1 5 9
3 1 10
`, []int64{49, 15, 23, 1, 9})
}

func TestModuloAndSet(t *testing.T) {
	runSample(t, `4 6
9 7 6 5
1 1 4
2 1 4 7
1 1 4
3 2 20
2 1 4 10
1 1 4
`, []int64{27, 13, 13})
}

func TestRandomDifferential(t *testing.T) {
	rng := rand.New(rand.NewSource(438))
	for it := 0; it < 200; it++ {
		n := rng.Intn(20) + 1
		a := make([]int, n)
		brute := make([]int, n)
		ops := make([][]int, 0)
		for i := range n {
			a[i] = rng.Intn(100) + 1
			brute[i] = a[i]
		}

		var want []int64
		for range 100 {
			typ := rng.Intn(3) + 1
			if typ == 1 {
				l := rng.Intn(n)
				r := rng.Intn(n-l) + l
				ops = append(ops, []int{1, l + 1, r + 1})
				var sum int64
				for i := l; i <= r; i++ {
					sum += int64(brute[i])
				}
				want = append(want, sum)
			} else if typ == 2 {
				l := rng.Intn(n)
				r := rng.Intn(n-l) + l
				x := rng.Intn(100) + 1
				ops = append(ops, []int{2, l + 1, r + 1, x})
				for i := l; i <= r; i++ {
					brute[i] %= x
				}
			} else {
				k := rng.Intn(n)
				x := rng.Intn(100) + 1
				ops = append(ops, []int{3, k + 1, x})
				brute[k] = x
			}
		}

		got := solve(a, ops)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("iteration %d: sample expect %v, but got %v", it, want, got)
		}
	}
}
