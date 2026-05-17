package p3933

import "testing"

func runSample(t *testing.T, mat [][]int, expect int) {
	res := countLocalMaximums(mat)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
	expect := 1
	runSample(t, mat, expect)
}

func TestSample2(t *testing.T) {
	mat := [][]int{
		{1, 2}, {3, 4},
	}
	expect := 1
	runSample(t, mat, expect)
}

func TestSample3(t *testing.T) {
	mat := [][]int{
		{1, 0, 1}, {0, 1, 0}, {1, 0, 1},
	}
	expect := 5
	runSample(t, mat, expect)
}

func TestSample4(t *testing.T) {
	mat := [][]int{
		{1, 1}, {1, 1},
	}
	expect := 4
	runSample(t, mat, expect)
}
