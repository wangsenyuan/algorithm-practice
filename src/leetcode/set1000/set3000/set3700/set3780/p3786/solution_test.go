package p3786

import "testing"

func runSample(t *testing.T, n int, edges [][]int, group []int, expect int64) {
	res := interactionCosts(n, edges, group)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{
		{0, 1},
		{1, 2},
	}
	group := []int{1, 1, 1}
	expect := int64(4)
	runSample(t, n, edges, group, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{
		{0, 1},
		{1, 2},
	}
	group := []int{3, 2, 3}
	expect := int64(2)
	runSample(t, n, edges, group, expect)
}
