package p3593

import "testing"

func runSample(t *testing.T, n int, edges [][]int, cost []int, expect int) {
	res := minIncrease(n, edges, cost)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1}, {0, 2}}
	cost := []int{2, 1, 3}
	expect := 1
	runSample(t, n, edges, cost, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}}
	cost := []int{5, 1, 4}
	expect := 0
	runSample(t, n, edges, cost, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	edges := [][]int{{0, 4}, {0, 1}, {1, 2}, {1, 3}}
	cost := []int{3, 4, 1, 1, 7}
	expect := 1
	runSample(t, n, edges, cost, expect)
}
