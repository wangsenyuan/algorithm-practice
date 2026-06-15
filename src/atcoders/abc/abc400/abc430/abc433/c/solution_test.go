package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "7788788"
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "1112222334445556555"
	expect := 11
	runSample(t, s, expect)
}
