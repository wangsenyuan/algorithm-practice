package p3855

import "testing"

func runSample(t *testing.T, l, r, k int, expect int) {
	res := sumOfNumbers(l, r, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	l := 1
	r := 2
	k := 2
	expect := 66
	runSample(t, l, r, k, expect)
}

func TestSample2(t *testing.T) {
	l := 0
	r := 1
	k := 3
	expect := 444
	runSample(t, l, r, k, expect)
}

func TestSample3(t *testing.T) {
	// 00 + 01 + 10 + 11 = 22
	l := 0
	r := 1
	k := 2
	expect := 22

	// 1 * pow(10, 0) * pow(2, 0) + 1 * pow(10, 1) * pow(2, 1) = 22
	// 1 + 20 = 22?	

	runSample(t, l, r, k, expect)
}
