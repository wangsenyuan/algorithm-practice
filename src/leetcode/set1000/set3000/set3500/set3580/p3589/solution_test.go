package p3589

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := primeSubarray(nums, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3}
	k := 1
	expect := 2
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 3, 5, 7}
	k := 3
	expect := 4
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{139, 26863, 4817}
	k := 24583
	expect := 1
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{35148, 33542, 31973}
	k := 8258
	expect := 0
	runSample(t, nums, k, expect)
}
