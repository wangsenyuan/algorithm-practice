package p3629

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minJumps(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 2, 4, 6}, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{2, 3, 4, 7, 9}, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{4, 6, 5, 8}, 3)
}
