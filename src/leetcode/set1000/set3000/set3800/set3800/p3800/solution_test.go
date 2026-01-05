package p3800

import "testing"

func runSample(t *testing.T, s, x string, flipCost, swapCost, crossCost int, expect int) {
	res := minimumCost(s, x, flipCost, swapCost, crossCost)
	if res != int64(expect) {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "01000"
	x := "10111"
	flipCost := 10
	swapCost := 2
	crossCost := 2
	expect := 16
	runSample(t, s, x, flipCost, swapCost, crossCost, expect)
}

func TestSample2(t *testing.T) {
	s := "001"
	x := "110"
	flipCost := 2
	swapCost := 100
	crossCost := 100
	expect := 6
	runSample(t, s, x, flipCost, swapCost, crossCost, expect)
}

func TestSample3(t *testing.T) {
	s := "1010"
	x := "1010"
	flipCost := 5
	swapCost := 5
	crossCost := 5
	expect := 0
	runSample(t, s, x, flipCost, swapCost, crossCost, expect)
}
