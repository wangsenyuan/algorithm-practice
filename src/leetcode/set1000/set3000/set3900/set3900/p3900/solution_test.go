package p3900

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := longestBalanced(s)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "100001"
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "111"
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "01110"
	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "01111100"
	expect := 6
	runSample(t, s, expect)
}