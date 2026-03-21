package main

import "testing"

func runSample(t *testing.T, a []int, unlucky []int, expect int) {
	t.Helper()
	res := solve(a, unlucky)
	if res != expect {
		t.Fatalf("expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{2, 3, 5}, []int{5, 7}, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{2, 2, 2}, []int{1, 3}, 6)
}

func TestAgainstBruteForce(t *testing.T) {
	cases := []struct {
		a       []int
		unlucky []int
	}{
		{[]int{1}, nil},
		{[]int{1, 2}, []int{1}},
		{[]int{1, 2}, []int{2}},
		{[]int{1, 2, 3}, []int{1, 3}},
		{[]int{1, 1, 2}, []int{2}},
		{[]int{2, 2, 3, 4}, []int{4, 7}},
	}

	for _, tc := range cases {
		expect := bruteForce(tc.a, tc.unlucky)
		runSample(t, tc.a, tc.unlucky, expect)
	}
}

func bruteForce(a []int, unlucky []int) int {
	bad := make(map[int]bool)
	for _, x := range unlucky {
		bad[x] = true
	}

	n := len(a)
	used := make([]bool, n)
	perm := make([]int, n)
	var ans int

	var dfs func(pos int)
	dfs = func(pos int) {
		if pos == n {
			sum := 0
			for _, idx := range perm {
				sum += a[idx]
				if bad[sum] {
					return
				}
			}
			ans++
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			perm[pos] = i
			dfs(pos + 1)
			used[i] = false
		}
	}

	dfs(0)
	return ans
}
