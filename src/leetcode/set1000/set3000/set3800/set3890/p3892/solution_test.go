package p3892

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := minOperations(nums, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 1, 2}
	k := 1
	expect := 1
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 5, 3, 6}
	k := 2
	expect := 0
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{3, 7, 3}
	k := 2
	expect := -1
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{1, -4}
	k := 0
	expect := 0
	runSample(t, nums, k, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{23, 10, -14, -23, 1}
	k := 2
	expect := 25
	runSample(t, nums, k, expect)
}
