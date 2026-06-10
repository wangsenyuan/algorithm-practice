package main

import "testing"

func runSample(t *testing.T, n int, m int, expect bool) {
	res := solve(n, m)

	if (len(res) == n) != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	// 怎么验证呢?
	freq := make(map[int]int)
	for i := 0; i < n; i++ {
		if i+1 < n && res[i] >= res[i+1] {
			t.Fatalf("Sample result %v, is invalid, it is not asc", res)
		}
		freq[res[i]]++
	}
	// 怎么验证m条件呢?
	// 比较小,直接brute froce

	var cnt int
	for i := range n {
		for j := i + 1; j < n; j++ {
			cnt += freq[res[i]+res[j]]
		}
	}

	if cnt != m {
		t.Fatalf("Sample result %v can't get %d tuples", res, m)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 3
	expect := true
	runSample(t, n, m, expect)
}

func TestSample2(t *testing.T) {
	n, m := 4, 10
	expect := false
	runSample(t, n, m, expect)
}
