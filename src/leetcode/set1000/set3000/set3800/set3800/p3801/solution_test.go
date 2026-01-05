package p3801

import "testing"

func runSample(t *testing.T, lists [][]int, expect int) {
	res := minMergeCost(lists)

	if res != int64(expect) {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	lists := [][]int{
		{1, 3, 5}, {2, 4}, {6, 7, 8},
	}
	expect := 18
	runSample(t, lists, expect)
}

func TestSample2(t *testing.T) {
	lists := [][]int{
		{1, 1, 5}, {1, 4, 7, 8},
	}
	expect := 10
	runSample(t, lists, expect)
}

func TestSample3(t *testing.T) {
	lists := [][]int{
		{3, 6, 10}, {-10, -5, -4, 1},
	}
	expect := 18
	runSample(t, lists, expect)
}

func TestSample4(t *testing.T) {
	lists := [][]int{
		{-7, 2}, {-9, -3, -3, 2, 5}, {-9, 2, 6, 6, 7}, {0, 2, 2, 4},
	}
	expect := 45
	runSample(t, lists, expect)
}
