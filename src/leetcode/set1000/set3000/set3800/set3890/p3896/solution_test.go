package p3896

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minOperations(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{4, 4}
	expect := 1
	runSample(t, nums, expect)
}
