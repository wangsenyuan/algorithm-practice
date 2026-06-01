package p3947

import "testing"

func runSample(t *testing.T, items [][]int, budget int, expect int) {
	res := maximumSaleItems(items, budget)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	items := [][]int{
		{1, 6}, {2, 4}, {3, 5},
	}
	budget := 19
	expect := 5
	runSample(t, items, budget, expect)
}

func TestSample2(t *testing.T) {
	items := [][]int{
		{2, 8}, {1, 10}, {6, 6}, {4, 12}, {5, 20}, {5, 17},
	}
	budget := 35
	expect := 7
	runSample(t, items, budget, expect)
}

func TestSample3(t *testing.T) {
	items := [][]int{
		{3, 8}, {3, 13}, {6, 28}, {6, 54}, {6, 34}, {6, 11}, {6, 19}, {6, 47}, {6, 14},
	}
	budget := 66
	expect := 16
	runSample(t, items, budget, expect)
}
