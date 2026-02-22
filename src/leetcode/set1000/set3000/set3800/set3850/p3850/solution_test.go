package p3850

import "testing"

func runSample(t *testing.T, nums []int, k int64, expect int) {
	res := countSequences(nums, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 3, 2}
	k := int64(6)
	expect := 2
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 6, 3}
	k := int64(2)
	expect := 2
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 5}
	k := int64(1)
	expect := 3
	runSample(t, nums, k, expect)
}
func TestSample4(t *testing.T) {
	nums := []int{5, 5}
	k := int64(1)
	expect := 3
	runSample(t, nums, k, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{5, 3, 5, 5}
	k := int64(15)
	expect := 6
	runSample(t, nums, k, expect)
}
