package p3654

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int64) {
	res := minArraySum(nums, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 1, 1}
	k := 2
	expect := int64(1)
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 1, 4, 1, 5}
	k := 3
	expect := int64(5)
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{36, 78, 29, 83, 81, 87, 45, 66, 63, 28}
	k := 93
	expect := int64(317)
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{58, 68, 57, 71, 52, 6, 40, 22, 13, 29, 26, 17, 47, 31, 51, 73, 59, 69, 37, 14}
	k := 34
	expect := int64(58)
	runSample(t, nums, k, expect)
}
