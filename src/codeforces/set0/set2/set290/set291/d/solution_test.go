package main

import "testing"

func runSample(t *testing.T, n int, k int) {
	res := solve(n, k)

	if len(res) != k {
		t.Fatalf("Sample expect %d, but got %d", k, len(res))
	}
	a := make([]int, n)
	for i := range n {
		a[i] = 1
	}
	a[n-1] = 0
	buf := make([]int, n)
	for _, cur := range res {
		if len(cur) != n {
			t.Fatalf("Sample result %v, not correct", res)
		}

		copy(buf, a)

		for i := range n {
			j := cur[i] - 1
			buf[i] += a[j]
		}

		copy(a, buf)
	}

	for i := range n {
		if a[i] != n-i-1 {
			t.Fatalf("Sample result %v, not correct => %v", res, a)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 1
	k := 1
	runSample(t, n, k)
}

func TestSample2(t *testing.T) {
	n := 3
	k := 2
	runSample(t, n, k)
}

func TestSample3(t *testing.T) {
	n := 100
	k := 10
	runSample(t, n, k)
}