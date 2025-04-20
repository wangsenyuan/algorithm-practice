package p3523

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, k int, expect []int64) {
	res := resultArray(nums, k)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := 3
	expect := []int64{9, 2, 4}
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 4, 8, 16, 32}
	k := 4
	expect := []int64{18, 1, 2, 0}
	runSample(t, nums, k, expect)
}
