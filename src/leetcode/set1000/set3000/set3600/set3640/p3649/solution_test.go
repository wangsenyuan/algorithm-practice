package p3649

import "testing"

func runSample(t *testing.T, arr []int, expect int64) {
	res := perfectPairs(arr)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{0, 1, 2, 3}
	expect := int64(2)
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{-3, 2, -1, 4}
	expect := int64(4)
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1, 10, 100, 1000}
	expect := int64(0)
	runSample(t, arr, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{0, 5, -2}
	expect := int64(0)
	runSample(t, arr, expect)
}
