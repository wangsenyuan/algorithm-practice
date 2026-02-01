package main

import "testing"


func runSample(t *testing.T, n int) {
	res := solve(n)
	if len(res) > 3 {
		t.Fatalf("Sample result %v not valid", res)
	}
	if res[0] + res[1] + res[2] != n {
		t.Fatalf("Sample result %v not valid", res)
	}
	if !checkPrime(res[0]) || !checkPrime(res[1]) || !checkPrime(res[2]) {
		t.Fatalf("Sample result %v not valid", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7865)
}
