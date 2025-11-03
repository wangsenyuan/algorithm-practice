package p3733

import "testing"

func runSample(t *testing.T, d []int, r []int, expect int) {
	res := minimumTime(d, r)
	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	d := []int{3, 1}
	r := []int{2, 3}
	expect := 5
	runSample(t, d, r, expect)
}

func TestSample2(t *testing.T) {
	d := []int{1, 3}
	r := []int{2, 2}
	expect := 7
	runSample(t, d, r, expect)
}

func TestSample3(t *testing.T) {
	d := []int{2, 1}
	r := []int{3, 4}
	expect := 3
	runSample(t, d, r, expect)
}
