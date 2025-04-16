package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "13"
	expect := "20"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "2345"
	expect := "2348"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "88"
	expect := "-1"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "1288"
	// 2 + 5 = 7
	expect := "1606"
	// 2 + 6 = 8
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "576079"
	expect := "576086"
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "26592659"
	expect := "26602660"
	runSample(t, s, expect)
}