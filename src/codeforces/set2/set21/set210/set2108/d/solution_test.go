package main

import "testing"

func runSample(t *testing.T, n int, k int, A []int, B []int, expect bool) {
	var cnt int
	ask := func(i int) int {
		cnt++
		if cnt > 250 {
			t.Fatalf("Sample asked too much %d", cnt)
		}
		if i <= len(A) {
			return A[i-1]
		}
		i -= len(A)
		return B[i-1]
	}

	res := solve(n, k, ask)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	if res[0] != len(A) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 2
	A := []int{1, 2}
	B := []int{2, 1, 2}
	expect := true
	runSample(t, n, k, A, B, expect)
}

func TestSample2(t *testing.T) {
	n, k := 18, 4
	A := []int{2, 4, 3, 1, 2, 4, 3, 1, 2}
	B := []int{1, 3, 2, 4, 1, 3, 2, 4, 1}
	expect := true
	runSample(t, n, k, A, B, expect)
}

func TestSample3(t *testing.T) {
	n, k := 3, 1
	A := []int{1, 1}
	B := []int{1}
	expect := false
	runSample(t, n, k, A, B, expect)
}

func TestSample4(t *testing.T) {
	n, k := 10, 5
	A := []int{1, 2, 3, 4, 5}
	B := []int{1, 2, 3, 4, 5}
	expect := true
	runSample(t, n, k, A, B, expect)
}

func TestSample5(t *testing.T) {
	n, k := 9, 3
	A := []int{1, 2, 3, 1, 2, 3}
	B := []int{1, 3, 2}
	expect := true
	runSample(t, n, k, A, B, expect)
}

func TestSample6(t *testing.T) {
	n, k := 12, 4
	A := []int{1, 3, 2, 4}
	B := []int{1, 3, 4, 2, 1, 3, 4, 2}
	expect := false
	runSample(t, n, k, A, B, expect)
}
