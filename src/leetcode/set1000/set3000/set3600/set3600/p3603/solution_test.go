package p3603

import "testing"

func runSample(t *testing.T, m int, n int, cost [][]int, expect int64) {
	res := minCost(m, n, cost)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 1
	n := 2
	waitCost := [][]int{
		{1, 2},
	}
	expect := int64(3)
	runSample(t, m, n, waitCost, expect)
}

func TestSample2(t *testing.T) {
	m := 2
	n := 2
	waitCost := [][]int{
		{3, 5},
		{2, 4},
	}
	expect := int64(9)
	runSample(t, m, n, waitCost, expect)
}

func TestSample3(t *testing.T) {
	m := 2
	n := 3
	waitCost := [][]int{
		{6, 1, 4},
		{3, 2, 5},
	}
	expect := int64(16)
	runSample(t, m, n, waitCost, expect)
}
