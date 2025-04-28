package p3530

import "testing"

func runSample(t *testing.T, n int, edges [][]int, score []int, expect int) {
	res := maxProfit(n, edges, score)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	edges := [][]int{
		{0, 1},
	}
	scores := []int{2, 3}
	expect := 8
	runSample(t, n, edges, scores, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{
		{0, 1}, {0, 2},
	}
	scores := []int{1, 6, 3}
	expect := 25
	runSample(t, n, edges, scores, expect)
}
