package p3575

import "testing"

func runSample(t *testing.T, vals []int, par []int, expect int) {
	ans := goodSubtreeSum(vals, par)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	vals := []int{2, 3}
	par := []int{-1, 0}
	expect := 8
	runSample(t, vals, par, expect)
}

func TestSample2(t *testing.T) {
	vals := []int{1, 5, 2}
	par := []int{-1, 0, 0}
	expect := 15
	runSample(t, vals, par, expect)
}

func TestSample3(t *testing.T) {
	vals := []int{34, 1, 2}
	par := []int{-1, 0, 1}
	expect := 42
	runSample(t, vals, par, expect)
}

func TestSample4(t *testing.T) {
	vals := []int{3, 22, 5}
	par := []int{-1, 0, 1}
	expect := 18
	runSample(t, vals, par, expect)
}

func TestSample5(t *testing.T) {
	vals := []int{5192, 4806, 1783, 6394, 7611}
	par := []int{-1, 3, 3, 0, 1}
	expect := 22981
	runSample(t, vals, par, expect)
}
