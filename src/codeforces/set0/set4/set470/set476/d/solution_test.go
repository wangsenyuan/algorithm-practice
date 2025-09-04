package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	min_num, res := solve(n, k)
	if min_num != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, min_num)
	}
	if len(res) != n {
		t.Fatalf("Sample result %v, not correct", res)
	}

	for _, cur := range res {
		for i := range 4 {
			for j := range i {
				if gcd(cur[i], cur[j]) != k {
					t.Fatalf("Sample result %v, not correct", cur)
				}
			}
		}
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 22)
}

func TestSample2(t *testing.T) {
	runSample(t, 7, 7, 287)
}
