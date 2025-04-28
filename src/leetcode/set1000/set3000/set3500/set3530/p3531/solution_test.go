package p3531

import (
	"testing"
)

func runSample(t *testing.T, n int, buildings [][]int, expect int) {
	res := countCoveredBuildings(n, buildings)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	buildings := [][]int{
		{1, 2}, {2, 2}, {3, 2}, {2, 1}, {2, 3},
	}
	expect := 1
	runSample(t, n, buildings, expect)
}
