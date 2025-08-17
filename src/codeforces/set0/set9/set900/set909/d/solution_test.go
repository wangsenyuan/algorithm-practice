package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "aabb", 2)
}

func TestSample2(t *testing.T) {
	runSample(t, "aabcaa", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, "aaaaaaaaabbbbbaaaabaaaaaaaaaaaaaaaaabaaaaaabbbbbbbaaabbbbbbbbbbbbbbbbbbbbbaaaaaaaaaaaaaaaaaaaaaaaaaa", 12)
}