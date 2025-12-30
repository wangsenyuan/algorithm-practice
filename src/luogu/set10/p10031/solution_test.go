package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3)
}

func TestSample2(t *testing.T) {
	// gcd(1, 6) ^ gcd(2, 6) ^ gcd(3, 6) ^ gcd(4, 6) ^ gcd(5, 6) ^ gcd(6, 6) = 5
	// 1 ^ 2^ 3 ^ 2 ^ 1 ^ 6 = 5
	runSample(t, 6, 5)
}
