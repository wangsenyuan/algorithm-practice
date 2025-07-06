package p3605

import "testing"

func runSample(t *testing.T, a []int, c int, expect int) {
	res := minStable(a, c)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 5, 10}
	maxC := 1
	expect := 1
	runSample(t, nums, maxC, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 6, 8}
	maxC := 1
	expect := 1
	runSample(t, nums, maxC, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 4, 9, 6}
	maxC := 1
	expect := 2
	runSample(t, nums, maxC, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{6, 5}
	maxC := 2
	expect := 0
	runSample(t, nums, maxC, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{169, 33, 168, 208, 184, 142}
	maxC := 0
	expect := 4
	runSample(t, nums, maxC, expect)
}
