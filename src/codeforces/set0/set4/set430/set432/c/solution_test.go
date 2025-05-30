package main

import (
	"math/rand/v2"
	"slices"
	"sort"
	"testing"
)

func runSample(t *testing.T, a []int) {
	n := len(a)
	arr := slices.Clone(a)
	res := solve(a)
	lpf := getPrimes(n)

	for _, op := range res {
		u, v := op[0], op[1]
		if u > v {
			u, v = v, u
		}
		d := v - u + 1
		if lpf[d] != d {
			t.Fatalf("Sample result %v, not a prime number %d", op, d)
		}
		arr[u-1], arr[v-1] = arr[v-1], arr[u-1]
	}

	if !sort.IntsAreSorted(arr) || len(res) > 5*n {
		t.Fatalf("Sample result %v, not sorted", arr)
	}
}

func TestSample1(t *testing.T) {
	a := []int{4, 2, 3, 1}
	runSample(t, a)
}

func TestSample2(t *testing.T) {
	n := 100
	a := make([]int, n)
	for i := range n {
		a[i] = i + 1
	}
	for range 3 {
		rand.Shuffle(n, func(i, j int) {
			a[i], a[j] = a[j], a[i]
		})
		runSample(t, a)
	}
}
