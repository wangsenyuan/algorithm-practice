package main

import "testing"

func runSample(t *testing.T, num string, expect string) {
	res := solve(num)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "527", "572")
}

func TestSample2(t *testing.T) {
	runSample(t, "4573", "3574")
}

func TestSample3(t *testing.T) {
	runSample(t, "1357997531", "-1")
}
