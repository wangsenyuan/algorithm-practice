package p3897

import "testing"

func runSample(t *testing.T, nums1 []int, nums0 []int, expect int) {
	res := maxValue(nums1, nums0)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums1 := []int{1, 2}
	nums0 := []int{1, 0}
	expect := 14
	runSample(t, nums1, nums0, expect)
}

func TestSample2(t *testing.T) {
	nums1 := []int{3, 1}
	nums0 := []int{0, 3}
	expect := 120
	runSample(t, nums1, nums0, expect)
}

func TestSample3(t *testing.T) {
	nums1 := []int{1, 2, 10000, 0}
	nums0 := []int{1, 0, 10, 3}
	expect := 349641308
	runSample(t, nums1, nums0, expect)
}
