package main

import "testing"

func runSample(t *testing.T, k int) {
	res := solve(k)

	var sum int
	freq := make([]int, 26)
	for _, ch := range res {
		x := int(ch - 'a')
		sum += freq[x]
		freq[x]++
	}
	if sum != k {
		t.Fatalf("Sample expect %d, but got %s(%d)", k, res, sum)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 12)
}
