package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "369727"
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "123456789987654321"
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "3636363636363454545454543636363636454545452727272727218181818181999111777"
	// 39  93
	// 49  94
	// 29 92
	// 19 91
	expect := 1512
	runSample(t, s, expect)
}
