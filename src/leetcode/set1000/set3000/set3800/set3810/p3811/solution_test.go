package p3811

import "testing"

func runSample(t *testing.T, nums []int, target1 int, target2 int, expect int) {
	res := alternatingXOR(nums, target1, target2)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 3, 1, 4}
	target1 := 1
	target2 := 5
	expect := 1
	runSample(t, nums, target1, target2, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 0, 0}
	target1 := 1
	target2 := 0
	expect := 3
	runSample(t, nums, target1, target2, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{7}
	target1 := 1
	target2 := 7
	expect := 0
	runSample(t, nums, target1, target2, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{17218, 0}
	target1 := 17218
	target2 := 27973
	expect := 1
	runSample(t, nums, target1, target2, expect)
}
