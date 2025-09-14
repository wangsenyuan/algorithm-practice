package p3685

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, k int, expect []bool) {
	res := subsequenceSumAfterCapping(nums, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{4, 3, 2, 4}
	k := 5
	expect := []bool{false, false, true, true}
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := 3
	expect := []bool{true, true, true, true, true}
	runSample(t, nums, k, expect)
}
