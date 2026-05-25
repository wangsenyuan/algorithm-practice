package p3939

import "testing"

func runSample(t *testing.T, parent []int, nums []int, k int, expect int) {
	res := countValidSubsets(parent, nums, k)
	if res != expect {
		t.Errorf("Sample %v %v %d, expect %d, but got %d", parent, nums, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	parent := []int{-1, 0, 1}
	nums := []int{1, 2, 3}
	k := 3
	expect := 1
	runSample(t, parent, nums, k, expect)
}

func TestSample2(t *testing.T) {
	parent := []int{-1, 0, 0, 0}
	nums := []int{2, 1, 2, 1}
	k := 3
	expect := 2
	runSample(t, parent, nums, k, expect)
}
