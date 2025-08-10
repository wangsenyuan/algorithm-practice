package p3645

import "testing"

func runSample(t *testing.T, value []int, limit []int, expect int) {
	res := maxTotal(value, limit)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	value := []int{3, 5, 8}
	limit := []int{2, 1, 3}
	expect := 16
	runSample(t, value, limit, expect)
}

func TestSample2(t *testing.T) {
	value := []int{4, 2, 6}
	limit := []int{1, 1, 1}
	expect := 6
	runSample(t, value, limit, expect)
}

func TestSample3(t *testing.T) {
	value := []int{4, 1, 5, 2}
	limit := []int{3, 3, 2, 3}
	expect := 12
	runSample(t, value, limit, expect)
}
