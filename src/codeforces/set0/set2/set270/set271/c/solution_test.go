package main

import "testing"

func runSample(t *testing.T, n int, k int, expect bool) {
	ans := solve(n, k)

	if len(ans) == n != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, ans)
	}
	if !expect {
		return
	}

	sets := make([][]int, k)
	for i, v := range ans {
		v--
		sets[v] = append(sets[v], i+1)
	}

	for _, cur := range sets {
		if len(cur) <= 2 {
			t.Fatalf("Sample result not correct, every group should have at least 3 elements, but got %v", cur)
		}
		d := cur[1] - cur[0]
		ok := false
		for j := 2; j < len(cur); j++ {
			if cur[j]-cur[j-1] != d {
				ok = true
				break
			}
		}
		if !ok {
			t.Fatalf("Sample result not correct, every group should not have arithmetic progression, but got %v", cur)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 11, 3, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 2, false)
}

func TestSample3(t *testing.T) {
	runSample(t, 9, 3, true)
}
