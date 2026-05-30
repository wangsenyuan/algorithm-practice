package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)
	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "0100"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "000"
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "11011011"
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "00100"
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "1"
	expect := true
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "01011"
	expect := true
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := "01"
	expect := true
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := "0101011"
	expect := true
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := "1101010"
	expect := true
	runSample(t, s, expect)
}

func TestSample10(t *testing.T) {
	s := "11001"
	expect := true
	runSample(t, s, expect)
}

func TestSample11(t *testing.T) {
	s := "1101"
	expect := false
	runSample(t, s, expect)
}

func TestSample12(t *testing.T) {
	s := "001101100"
	expect := false
	runSample(t, s, expect)
}
