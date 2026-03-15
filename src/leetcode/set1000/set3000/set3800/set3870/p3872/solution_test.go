package p3872

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := longestArithmetic(nums)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{90033, 1535, 13037, 24539, 842}
	expect := 4
	runSample(t, nums, expect)
}
