package main

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := solve(k, s)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "4727447"
	k := 4
	expect := "4427477"
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "4211147"
	k := 7
	expect := "4211177"
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "4747477"
	k := 6
	expect := "4444477"
	runSample(t, s, k, expect)
}
