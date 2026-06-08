package p3953

import "testing"

func runSample(t *testing.T, nums []int, mavValue int, expect int) {
	res := maxScore(nums, mavValue)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 4, 6}
	maxValue := 5
	expect := 4
	runSample(t, nums, maxValue, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3}
	maxValue := 4
	expect := 3
	runSample(t, nums, maxValue, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 2}
	maxValue := 1
	expect := 1
	runSample(t, nums, maxValue, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{42}
	maxValue := 312
	expect := 311
	runSample(t, nums, maxValue, expect)
}
