package p3757

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := countEffective(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{7, 4, 6}
	expect := 4
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{8, 8}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{2, 2, 1}
	expect := 5
	runSample(t, nums, expect)
}
