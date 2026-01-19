package p3814

import "testing"

func runSample(t *testing.T, costs []int, capacity []int, budget int, expect int) {
	res := maxCapacity(costs, capacity, budget)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	costs := []int{4, 8, 5, 3}
	capacity := []int{1, 5, 2, 7}
	budget := 8
	expect := 8
	runSample(t, costs, capacity, budget, expect)
}
