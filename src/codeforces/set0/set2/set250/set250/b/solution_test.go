package main

import "testing"

func runSample(t *testing.T, s string, expect string)  {
	res := solve(s)
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `a56f:d3:0:0124:01:f19a:1000:00`
	expect := "a56f:00d3:0000:0124:0001:f19a:1000:0000"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `a56f:00d3:0000:0124:0001::`
	expect := "a56f:00d3:0000:0124:0001:0000:0000:0000"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `a56f::0124:0001:0000:1234:0ff0`
	expect := "a56f:0000:0000:0124:0001:0000:1234:0ff0"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `a56f:0000::0000:0001:0000:1234:0ff0`
	expect := "a56f:0000:0000:0000:0001:0000:1234:0ff0"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `a56f:0000::0000:0001:0000:1234:0ff0`
	expect := "a56f:0000:0000:0000:0001:0000:1234:0ff0"
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `::`
	expect := "0000:0000:0000:0000:0000:0000:0000:0000"
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `0ea::4d:f4:6:0`
	expect := "00ea:0000:0000:0000:004d:00f4:0006:0000"
	runSample(t, s, expect)
}