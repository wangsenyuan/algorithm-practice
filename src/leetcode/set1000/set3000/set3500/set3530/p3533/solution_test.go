package p3533

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, k int, expect []int) {
	res := concatenatedDivisibility(nums, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 12, 45}
	k := 5
	expect := []int{3, 12, 45}
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{10, 5}
	k := 5
	expect := []int{5, 10}
	runSample(t, nums, k, expect)
}
