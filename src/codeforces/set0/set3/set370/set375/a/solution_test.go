package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, s string) {
	res := solve(s)

	getFreq := func(s string) []int {
		freq := make([]int, 10)
		for _, c := range s {
			freq[int(c-'0')]++
		}
		return freq
	}
	f1 := getFreq(s)
	f2 := getFreq(res)
	if !slices.Equal(f1, f2) {
		t.Fatalf("Sample expect %s, but got %s", s, res)
	}
	if res[0] == '0' {
		t.Fatalf("Sample result %s, leading zero", res)
	}
	var sum int
	for _, c := range res {
		sum = sum*10 + int(c-'0')
		sum %= 7
	}
	if sum != 0 {
		t.Fatalf("Sample result %s, got wrong sum %d", res, sum)
	}
}

func TestSample1(t *testing.T) {
	s := "1689"
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := "18906"
	runSample(t, s)
}
