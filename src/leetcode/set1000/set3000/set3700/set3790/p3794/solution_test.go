package p3794

import "testing"

func runSample(t *testing.T, n int, restrictions [][]int, diff []int, expect int) {
	res := findMaxVal(n, restrictions, diff)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	restrictions := [][]int{
		{3, 1}, {8, 1},
	}
	diff := []int{2, 2, 3, 1, 4, 5, 1, 1, 2}
	expect := 6
	runSample(t, n, restrictions, diff, expect)
}

func TestSample2(t *testing.T) {
	n := 8
	restrictions := [][]int{
		{3, 2},
	}
	diff := []int{3, 5, 2, 4, 2, 3, 1}
	expect := 12
	runSample(t, n, restrictions, diff, expect)
}
