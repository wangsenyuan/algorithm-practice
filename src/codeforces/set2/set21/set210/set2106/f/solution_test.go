package main

import "testing"

func runSample(t *testing.T, a string, expect int) {
	res := solve(a)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "000"
	expect := 3
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := "0010"
	expect := 9
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := "1011001"
	expect := 10
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := "0001"
	expect := 7
	runSample(t, a, expect)
}


func TestSample5(t *testing.T) {
	a := "11"
	expect := 1
	runSample(t, a, expect)
}