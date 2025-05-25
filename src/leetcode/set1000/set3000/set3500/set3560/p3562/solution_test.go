package p3562

import "testing"

func runSample(t *testing.T, n int, present []int, future []int, hierarchy [][]int, budget int, expect int) {
	res := maxProfit(n, present, future, hierarchy, budget)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	present := []int{1, 2}
	future := []int{4, 3}
	hierarchy := [][]int{{1, 2}}
	budget := 3
	expect := 5
	runSample(t, n, present, future, hierarchy, budget, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	present := []int{3, 4}
	future := []int{5, 8}
	hierarchy := [][]int{{1, 2}}
	budget := 4
	expect := 4
	runSample(t, n, present, future, hierarchy, budget, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	present := []int{4, 6, 8}
	future := []int{7, 9, 11}
	hierarchy := [][]int{{1, 2}, {1, 3}}
	budget := 10
	expect := 10
	runSample(t, n, present, future, hierarchy, budget, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	present := []int{5, 2, 3}
	future := []int{8, 5, 6}
	hierarchy := [][]int{{1, 2}, {2, 3}}
	budget := 7
	expect := 12
	runSample(t, n, present, future, hierarchy, budget, expect)
}
