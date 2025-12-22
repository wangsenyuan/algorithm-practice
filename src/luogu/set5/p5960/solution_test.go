package main

import "testing"

func runSample(t *testing.T, n int, checks [][]int, expect bool) {
	res := solve(n, checks)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	for _, ck := range checks {
		i, j, c := ck[0]-1, ck[1]-1, ck[2]
		if res[i]-res[j] > c {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 3
	checks := [][]int{
		{1, 2, 3},
		{2, 3, -2},
		{1, 3, 1},
	}

	expect := true
	runSample(t, n, checks, expect)
}
