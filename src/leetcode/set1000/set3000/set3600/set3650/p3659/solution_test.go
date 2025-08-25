package p3659

import "testing"

func runSample(t *testing.T, nums []int, k int, expect bool) {
	res := partitionArray(nums, k)
	if res != expect {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 5, 2, 2}
	k := 2
	expect := true
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{37, 96, 96, 59}
	k := 3
	expect := false
	runSample(t, nums, k, expect)
}
