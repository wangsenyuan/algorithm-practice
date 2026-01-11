package p3806

import (
	"testing"
)

func runSample(t *testing.T, nums []int, k int, m int, expect int) {
	res := maximumAND(nums, k, m)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 1, 2}
	k := 8
	m := 2
	expect := 6
	runSample(t, nums, k, m, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 8, 4}
	k := 7
	m := 3
	expect := 4
	runSample(t, nums, k, m, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 1}
	k := 3
	m := 2
	expect := 2
	runSample(t, nums, k, m, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{1}
	k := 4
	m := 1
	expect := 5
	runSample(t, nums, k, m, expect)
}
