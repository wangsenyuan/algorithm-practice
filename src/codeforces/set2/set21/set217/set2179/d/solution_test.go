package main

import (
	"math/bits"
	"testing"
)

func runSample(t *testing.T, n int, expect []int) {
	res := solve(n)

	get := func(arr []int) int {
		if len(arr) != 1<<n {
			t.Fatalf("Sample result %v, length not correct", arr)
		}
		var res int
		w := (1 << n) - 1
		vis := make([]bool, 1<<n)
		for _, v := range arr {
			if vis[v] {
				t.Fatalf("Sample result %v, not correct", arr)
			}
			vis[v] = true
			w &= v
			res += bits.OnesCount(uint(w))
		}

		return res
	}

	x := get(expect)
	y := get(res)

	if x != y {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, []int{1, 0})
}

func TestSample2(t *testing.T) {
	runSample(t, 2, []int{3, 1, 0, 2})
}

func TestSample3(t *testing.T) {
	runSample(t, 3, []int{7, 3, 1, 5, 0, 2, 4, 6})
}
