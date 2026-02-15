package p3845

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := maxXor(nums, k)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{5, 4, 5, 6}
	k := 2
	expect := 7
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 4, 5, 6}
	k := 1
	expect := 6
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{32767}
	k := 0
	expect := 32767
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{3, 4}
	k := 1
	expect := 7
	runSample(t, nums, k, expect)
}
