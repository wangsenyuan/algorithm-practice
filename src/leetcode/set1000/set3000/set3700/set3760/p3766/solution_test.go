package p3766

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int) {
	res := minOperations(nums)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 4}
	expect := []int{0, 1, 1}
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{6, 7, 12}
	expect := []int{1, 0, 3}
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2851}
	expect := []int{22}
	runSample(t, nums, expect)
}
