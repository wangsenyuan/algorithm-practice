package p3639

import "testing"

func runSample(t *testing.T, s string, order []int, k int, expect int) {
	res := minTime(s, order, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abc"
	order := []int{1, 0, 2}
	k := 2
	expect := 0
	runSample(t, s, order, k, expect)
}

func TestSample2(t *testing.T) {
	s := "cat"
	order := []int{0, 2, 1}
	k := 6
	expect := 2
	runSample(t, s, order, k, expect)
}

func TestSample3(t *testing.T) {
	s := "xy"
	order := []int{0, 1}
	k := 4
	expect := -1
	runSample(t, s, order, k, expect)
}
