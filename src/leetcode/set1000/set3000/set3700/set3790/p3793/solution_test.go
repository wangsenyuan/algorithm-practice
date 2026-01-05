package p3793

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := minLength(nums, k)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{6}
	k := 3
	expect := 1
	runSample(t, nums, k, expect)
}
