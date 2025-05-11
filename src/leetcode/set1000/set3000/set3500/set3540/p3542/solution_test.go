package p3542

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minOperations(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{0, 2}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 1, 2, 1}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 1, 2, 1, 2}
	expect := 4
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{0}
	expect := 0
	runSample(t, nums, expect)
}
