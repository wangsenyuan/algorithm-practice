package main

import "testing"

func runSample(t *testing.T, n int, d int, l int, expect bool) {
	res := solve(n, d, l)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	if !expect {
		return
	}

	sum := make([]int, 2)
	for i := range n {
		if i%2 == 0 {
			sum[0] += res[i]
		} else {
			sum[1] += res[i]
		}
		if res[i] < 1 || res[i] > l {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
	if sum[0]-sum[1] != d {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 3, 2, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, -4, 3, false)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, -4, 4, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 10, 5, 3, true)
}
