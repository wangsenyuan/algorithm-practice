package p3690

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int64) {
	res := maxTotalValue(nums, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{11, 8}
	k := 2
	expect := int64(6)
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{11, 8}
	k := 3
	expect := int64(9)
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 3, 2}
	k := 2
	expect := int64(4)
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{4, 2, 5, 1}
	k := 3
	expect := int64(12)
	runSample(t, nums, k, expect)
}
