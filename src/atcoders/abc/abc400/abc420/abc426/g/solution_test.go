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
	runSample(t, `4
3 4
5 8
1 2
2 3
3
1 4 7
2 4 10
1 2 2
`, []int{11, 13, 0})
}

func TestSample2(t *testing.T) {
	runSample(t, `8
167 430302156
22 623690081
197 476190629
176 24979445
22 877914575
247 211047202
232 822804784
25 628894325
8
6 8 176
3 5 80
1 7 310
4 8 368
4 5 218
3 4 431
4 6 228
1 1 239
`, []int{628894325, 877914575, 2324409440, 2329613684, 902894020, 501170074, 902894020, 430302156})
}

func solveBruteForce(items [][]int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for id, query := range queries {
		l, r, capacity := query[0]-1, query[1], query[2]
		dp := make([]int, capacity+1)
		for i := l; i < r; i++ {
			weight, value := items[i][0], items[i][1]
			for c := capacity; c >= weight; c-- {
				dp[c] = max(dp[c], dp[c-weight]+value)
			}
		}
		ans[id] = dp[capacity]
	}
	return ans
}

func TestSolveAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(426))
	for n := 1; n <= 20; n++ {
		items := make([][]int, n)
		for i := range items {
			items[i] = []int{rng.Intn(500) + 1, rng.Intn(1_000_000_000) + 1}
		}

		queries := make([][]int, 100)
		for i := range queries {
			l := rng.Intn(n)
			r := rng.Intn(n-l) + l + 1
			queries[i] = []int{l + 1, r, rng.Intn(500) + 1}
		}

		expect := solveBruteForce(items, queries)
		if result := solve(items, queries); !reflect.DeepEqual(result, expect) {
			t.Fatalf("n = %d, expect %v, but got %v", n, expect, result)
		}
	}
}

func TestLargeNumberOfQueries(t *testing.T) {
	const (
		n = 20000
		q = 200000
	)
	items := make([][]int, n)
	for i := range items {
		items[i] = []int{1, 1}
	}
	queries := make([][]int, q)
	for i := range queries {
		l := i * 7919 % n
		length := i*104729%(n-l) + 1
		queries[i] = []int{l + 1, l + length, i%500 + 1}
	}

	result := solve(items, queries)
	for i, value := range result {
		expect := min(queries[i][1]-queries[i][0]+1, queries[i][2])
		if value != expect {
			t.Fatalf("query %d, expect %d, but got %d", i, expect, value)
		}
	}
}
