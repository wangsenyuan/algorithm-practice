package p3960

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := getLength(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 2, 1, 2, 3, 3, 3}
	expect := 5
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expect := 1
	runSample(t, nums, expect)
}
