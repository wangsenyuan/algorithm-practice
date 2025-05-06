package p3538

import (
	"testing"
)

func runSample(t *testing.T, l int, n int, k int, position []int, time []int, expect int) {
	res := minTravelTime(l, n, k, position, time)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	l := 10
	n := 4
	k := 1
	position := []int{0, 3, 8, 10}
	time := []int{5, 8, 3, 6}
	expect := 62
	runSample(t, l, n, k, position, time, expect)
}

func TestSample2(t *testing.T) {
	l := 5
	n := 5
	k := 1
	position := []int{0, 1, 2, 3, 5}
	time := []int{8, 3, 9, 3, 3}
	expect := 34
	runSample(t, l, n, k, position, time, expect)
}

func TestSample3(t *testing.T) {
	l := 3
	n := 3
	k := 1
	position := []int{0, 1, 3}
	time := []int{1, 3, 1}
	expect := 3
	runSample(t, l, n, k, position, time, expect)
}
