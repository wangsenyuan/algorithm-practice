package p3661

import "testing"

func runSample(t *testing.T, robots []int, distance []int, walls []int, expect int) {
	res := maxWalls(robots, distance, walls)
	if res != expect {
		t.Errorf("Sample %v %v %v, expect %d, but got %d", robots, distance, walls, expect, res)
	}
}

func TestSample1(t *testing.T) {
	robots := []int{4}
	distance := []int{3}
	walls := []int{1, 10}
	runSample(t, robots, distance, walls, 1)
}

func TestSample2(t *testing.T) {
	robots := []int{10, 2}
	distance := []int{5, 1}
	walls := []int{5, 2, 7}
	runSample(t, robots, distance, walls, 3)
}

func TestSample3(t *testing.T) {
	robots := []int{1, 2}
	distance := []int{100, 1}
	walls := []int{10}
	runSample(t, robots, distance, walls, 0)
}

func TestSample4(t *testing.T) {
	robots := []int{63, 56, 40, 45, 4, 9, 44, 69, 55, 26, 73, 15, 12, 60, 43, 39, 37, 74, 36, 34, 13, 23, 66, 14, 11, 42, 72, 3, 57, 10, 53, 8, 70, 17, 58, 61, 30, 32}
	distance := []int{8, 7, 4, 8, 9, 5, 2, 4, 5, 2, 6, 9, 5, 9, 5, 3, 7, 6, 9, 2, 8, 7, 4, 3, 5, 1, 7, 5, 1, 3, 5, 3, 5, 4, 8, 7, 6, 4}
	walls := []int{6, 22, 50, 52, 20, 9, 23, 75, 26, 21, 60, 58, 41, 28, 30}
	runSample(t, robots, distance, walls, 15)
}
