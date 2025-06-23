package p3587

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minSwaps(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 4, 6, 5, 7}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 4, 5, 7}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 3}
	expect := 0
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{31, 248, 172}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{16, 540, 674, 760, 358, 463, 5, 445, 549}
	expect := 10
	runSample(t, nums, expect)
}
