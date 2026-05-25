package p3937

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := minOperations(nums, k)
	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", nums, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 4, 2, 8}
	k := 3
	expect := 2
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 1, 1}
	k := 3
	expect := 1
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{94, 53, 84}
	k := 20
	expect := 10
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{78, 339182, 56, 83568}
	k := 8
	expect := 4
	runSample(t, nums, k, expect)
}
