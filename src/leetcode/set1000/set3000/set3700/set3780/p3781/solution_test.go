package p3781

import "testing"

func runSample(t *testing.T, nums []int, s string, expect int64) {
	res := maximumScore(nums, s)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 1, 5, 2, 3}
	s := "01010"
	expect := int64(7)
	runSample(t, nums, s, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 7, 2, 9}
	s := "0000"
	expect := int64(0)
	runSample(t, nums, s, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{7, 9, 4, 2}
	s := "0011"
	expect := int64(16)
	runSample(t, nums, s, expect)
}
