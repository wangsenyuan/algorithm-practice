package p3694

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := distinctPoints(s, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "UD"
	k := 1
	expect := 2
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "UDLR"
	k := 4
	expect := 1
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "UU"
	k := 1
	expect := 1
	runSample(t, s, k, expect)
}
