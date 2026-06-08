package p3952

import "testing"

func runSample(t *testing.T, nums []int, s string, expect int) {
	res := maxTotal(nums, s)

	if int64(expect) != res {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{9, 2, 6, 1}
	s := "0101"
	expect := 15
	runSample(t, nums, s, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 1, 4}
	s := "001"
	expect := 4
	runSample(t, nums, s, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{9, 3, 5}
	s := "011"
	expect := 14
	runSample(t, nums, s, expect)
}
