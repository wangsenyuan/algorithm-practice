package p3584

import "testing"

func runSample(t *testing.T, nums []int, m int, expect int64) {
	res := maximumProduct(nums, m)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{0, -4, 1, 0, 1}
	m := 3
	expect := int64(1)
	runSample(t, nums, m, expect)
}
