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
	runSample(t, `1
5 3
2 3 5 7 11
2 1 5
2 2 4
2310 1 5
`, []int{5, 6, 1629})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
4 3
18 12 8 9
216 1 2
48 2 4
82944 1 4
`, []int{13, 12, 520})
}

func TestLargeDuplicateValues(t *testing.T) {
	const n = 100000
	a := make([]int, n)
	for i := range a {
		a[i] = 2
	}

	const q = 50000
	queries := make([][]int, q)
	expect := make([]int, q)
	for i := range q {
		queries[i] = []int{100000, 1, n}
		expect[i] = 312500000
	}
	res := solve(a, queries)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("expect every answer to be %d", expect[0])
	}
}

func TestRandomAgainstBrute(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for tc := 0; tc < 1000; tc++ {
		n := rng.Intn(8) + 1
		a := make([]int, n)
		for i := range a {
			a[i] = rng.Intn(11) + 2
		}

		q := rng.Intn(10) + 1
		queries := make([][]int, q)
		expect := make([]int, q)
		for i := range q {
			l := rng.Intn(n)
			r := rng.Intn(n-l) + l
			k := rng.Intn(100) + 1
			queries[i] = []int{k, l + 1, r + 1}
			for j := l; j <= r; j++ {
				for k%a[j] == 0 {
					k /= a[j]
				}
				expect[i] += k
			}
		}

		res := solve(a, queries)
		if !reflect.DeepEqual(res, expect) {
			t.Fatalf("a=%v, queries=%v, expect %v, but got %v", a, queries, expect, res)
		}
	}
}
