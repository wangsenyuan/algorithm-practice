package main

import "testing"

func runSample(t *testing.T, x string, expect string) {
	res := solve(x)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "16"
	expect := "1.6E1"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "01.23400"
	expect := "1.234"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := ".100"
	expect := "1E-1"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "100"
	expect := "1E2"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "0.0012"
	expect := "1.2E-3"
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "143529100720960530144687499862369157252883621496987867683546098241081752607457981824764693332677189."
	expect := "1.43529100720960530144687499862369157252883621496987867683546098241081752607457981824764693332677189E98"
	runSample(t, s, expect)
}
