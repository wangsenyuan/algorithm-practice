package p3961

import "testing"

func runSample(t *testing.T, units [][]int, expect int) {
	res := maxRatings(units)

	if int64(expect) != res {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	units := [][]int{
		{1, 3}, {2, 2},
	}
	expect := 4
	runSample(t, units, expect)
}

func TestSample2(t *testing.T) {
	units := [][]int{
		{1, 2, 3}, {4, 5, 6},
	}
	expect := 6
	runSample(t, units, expect)
}

func TestSample3(t *testing.T) {
	units := [][]int{
		{1, 1, 1}, {5, 5, 5},
	}
	expect := 6
	runSample(t, units, expect)
}
