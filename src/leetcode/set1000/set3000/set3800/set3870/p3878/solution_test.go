package p3878

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := countGoodSubarrays(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{4, 2, 3}
	expect := int64(4)
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 3, 1}
	expect := int64(6)
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{6, 6}
	expect := int64(3)
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{6, 6, 6}
	// 3 + 1 + 1 + 1 = 6
	expect := int64(6)
	runSample(t, nums, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{5, 2, 7}
	// 5, 2, 7
	expect := int64(5)
	runSample(t, nums, expect)
}