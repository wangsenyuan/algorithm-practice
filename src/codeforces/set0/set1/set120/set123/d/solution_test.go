package main

import (
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aaaa"
	expect := 20
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "abcdef"
	expect := 21
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "abacabadabacaba"
	expect := 188
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "jsoxkutcvyshsinfmtrpujedcbmyqlojzco"
	expect := 646
	runSample(t, s, expect)
}
