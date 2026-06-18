package main

import (
	"testing"
)

func runSample(t *testing.T, n int, expect []int) {
	t.Helper()

	marked := make([]bool, n+1)

	check := func(arr []int) int {
		clear(marked)
		var res int
		for i, v := range arr {
			if gcd(i+1, v) == 1 && i > 0 {
				t.Fatalf("Sample result %v is not valid", arr)
			}
			if i+1 == v {
				res++
			}
			if marked[v] {
				t.Fatalf("Sample result has duplicates %d", v)
			}
		}
		return res
	}

	res := solve(n)

	x := check(expect)
	y := check(res)

	if x != y {
		t.Fatalf("Sample result %v, is wrong", res)
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func TestSample1(t *testing.T) {
	n := 2
	expect := []int{1, 2}
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	expect := []int{1, 4, 6, 2, 5, 3}
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	n := 13
	expect := []int{1, 12, 9, 6, 10, 8, 7, 4, 3, 5, 11, 2, 13}
	runSample(t, n, expect)
}
