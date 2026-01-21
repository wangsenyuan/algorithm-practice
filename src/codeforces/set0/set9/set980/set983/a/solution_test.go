package main

import (
	"testing"
)


func runSample(t *testing.T, p int, q int, b int, expect string) {
	res := solve(p, q, b)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}
func TestSample1(t *testing.T) {
	p, q, b := 6, 12, 10
	expect := "Finite"
	runSample(t, p, q, b, expect)
}