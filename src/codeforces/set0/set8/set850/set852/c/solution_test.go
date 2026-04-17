package main

import "testing"

func score(court []int, bill []int) int {
	n := len(court)
	var res int
	for i := 0; i < n; i++ {
		res += bill[i] * (court[i] + court[(i+1)%n])
	}
	return res
}

func isPerm(a []int) bool {
	n := len(a)
	seen := make([]bool, n)
	for _, x := range a {
		if x < 0 || x >= n || seen[x] {
			return false
		}
		seen[x] = true
	}
	return true
}

func brute(court []int) []int {
	n := len(court)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	best := append([]int(nil), p...)
	bestScore := score(court, p)

	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			cur := score(court, p)
			if cur > bestScore {
				bestScore = cur
				copy(best, p)
			}
			return
		}
		for j := i; j < n; j++ {
			p[i], p[j] = p[j], p[i]
			dfs(i + 1)
			p[i], p[j] = p[j], p[i]
		}
	}

	dfs(0)
	return best
}

func runSample(t *testing.T, court []int) {
	t.Helper()
	res := solve(court)
	if !isPerm(res) {
		t.Fatalf("result is not a permutation: %v", res)
	}
	if len(court) <= 8 {
		expect := brute(court)
		if score(court, res) != score(court, expect) {
			t.Fatalf("expect optimal score %d, got %d, res=%v", score(court, expect), score(court, res), res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{0, 1, 2})
}

func TestSmall1(t *testing.T) {
	runSample(t, []int{1, 0})
}

func TestSmall2(t *testing.T) {
	runSample(t, []int{2, 0, 1})
}

func TestSmall3(t *testing.T) {
	runSample(t, []int{3, 1, 0, 2})
}
