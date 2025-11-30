package p3761

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minMirrorPairDistance(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{120, 21}
	expect := 1
	runSample(t, nums, expect)
}
