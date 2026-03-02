package p3859

import "testing"

func runSample(t *testing.T, nums []int, k int, m int, expect int64) {
	res := countSubarrays(nums, k, m)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 1, 2, 2}
	k := 2
	m := 2
	expect := int64(2)
	runSample(t, nums, k, m, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 1, 2, 4}
	k := 2
	m := 1
	expect := int64(3)
	runSample(t, nums, k, m, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 1, 2}
	k := 1
	m := 2
	expect := int64(1)
	runSample(t, nums, k, m, expect)
}
