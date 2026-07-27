package main

import (
	"bufio"
	"math/rand"
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
	runSample(t, `6
2 6 4 5 7 5
5
1 6 1
3 5 4
4 4 1
2 5 1
1 6 100
`, []int{6, 11, 0, 2, 10})
}

func TestFullRangeThenSubrange(t *testing.T) {
	runSample(t, `2
5 5
2
1 2 1
1 1 1
`, []int{2, 1})
}

func TestSubrangeThenFullRange(t *testing.T) {
	runSample(t, `2
5 5
2
1 1 1
1 2 5
`, []int{1, 9})
}

func TestRandomAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(1))

	for tc := 0; tc < 1000; tc++ {
		n := rng.Intn(8) + 1
		a := make([]int, n)
		stock := make([]int, n)
		for i := range n {
			a[i] = rng.Intn(20) + 1
			stock[i] = a[i]
		}

		q := rng.Intn(30) + 1
		queries := make([][]int, q)
		expect := make([]int, q)
		for i := range q {
			l := rng.Intn(n)
			r := rng.Intn(n-l) + l
			k := rng.Intn(10) + 1
			queries[i] = []int{l + 1, r + 1, k}

			for j := l; j <= r; j++ {
				buy := min(stock[j], k)
				expect[i] += buy
				stock[j] -= buy
			}
		}

		if got := solve(a, queries); !reflect.DeepEqual(got, expect) {
			t.Fatalf("case %d: a=%v queries=%v, expect %v, got %v",
				tc, a, queries, expect, got)
		}
	}
}
