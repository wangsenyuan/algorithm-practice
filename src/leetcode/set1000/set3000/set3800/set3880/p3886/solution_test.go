package p3886

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := sortableIntegers(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 1, 2}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{7, 6, 5}
	expect := 0
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{5, 8}
	expect := 3
	runSample(t, nums, expect)
}
