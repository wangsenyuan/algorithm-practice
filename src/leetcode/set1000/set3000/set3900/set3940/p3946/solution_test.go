package p3946

import (
	"testing"
)

func runSample(t *testing.T, items [][]int, budget int, expect int) {
	res := maximumSaleItems(items, budget)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	items := [][]int{
		{6, 2}, {2, 6}, {3, 4},
	}
	budget := 9
	expect := 4
	runSample(t, items, budget, expect)
}

func TestSample2(t *testing.T) {
	items := [][]int{
		{2, 4}, {3, 2}, {4, 1}, {6, 4}, {12, 4},
	}
	budget := 8
	expect := 10
	runSample(t, items, budget, expect)
}
