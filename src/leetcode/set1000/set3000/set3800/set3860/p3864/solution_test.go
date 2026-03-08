package p3864

import "testing"

func runSample(t *testing.T, s string, encCost int, flatCost int, expect int64) {
	res := minCost(s, encCost, flatCost)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1010"
	encCost := 2
	flatCost := 1
	expect := int64(6)
	runSample(t, s, encCost, flatCost, expect)
}

func TestSample2(t *testing.T) {
	s := "1010"
	encCost := 3
	flatCost := 10
	expect := int64(12)
	runSample(t, s, encCost, flatCost, expect)
}

func TestSample3(t *testing.T) {
	s := "00"
	encCost := 1
	flatCost := 2
	expect := int64(2)
	runSample(t, s, encCost, flatCost, expect)
}

func TestSample4(t *testing.T) {
	s := "010"
	encCost := 22
	flatCost := 8
	expect := int64(66)
	runSample(t, s, encCost, flatCost, expect)
}
