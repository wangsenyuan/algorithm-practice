package p3660

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int) {
	res := maxValue(nums)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 1, 3}
	expect := []int{2, 2, 3}
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 3, 1}
	expect := []int{3, 3, 3}
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{20, 21, 25, 15}
	expect := []int{25, 25, 25, 25}
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{19, 25, 12, 21}
	expect := []int{25, 25, 25, 25}
	runSample(t, nums, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{30, 21, 5, 35, 24}
	expect := []int{35, 35, 35, 35, 35}
	runSample(t, nums, expect)
}

func TestSample6(t *testing.T) {
	nums := []int{7, 8, 6, 1}
	expect := []int{8, 8, 8, 8}
	runSample(t, nums, expect)
}
