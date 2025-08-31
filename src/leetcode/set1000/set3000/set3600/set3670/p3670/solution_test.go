package p3670

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxProduct(nums)
	if res != int64(expect) {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	expect := 12
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 6, 4}
	expect := 0
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{64, 8, 32}
	expect := 2048
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{7, 17, 7}
	expect := 0
	runSample(t, nums, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{9, 2, 19}
	expect := 18
	runSample(t, nums, expect)
}
