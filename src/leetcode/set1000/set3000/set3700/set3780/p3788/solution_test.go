package p3788

import "testing"


func runSample(t *testing.T, nums []int, expect int) {
	res := maximumScore(nums)
	if res != int64(expect) {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{10,-1,3,-4,-5}
	expect := 17
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{-7,-5,3}
	expect := -2
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1,1}
	expect := 0
	runSample(t, nums, expect)
}