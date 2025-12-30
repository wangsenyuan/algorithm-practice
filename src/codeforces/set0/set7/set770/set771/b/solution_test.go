package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, k, x, res := drive(reader)

	freq := make(map[string]int)

	for i := range n {
		freq[res[i]]++
		if i >= k {
			freq[res[i-k]]--
			if freq[res[i-k]] == 0 {
				delete(freq, res[i-k])
			}
		}
		if i >= k-1 {
			if (len(freq) == k) != (x[i-k+1] == "YES") {
				t.Fatalf("Sample expect %s at %d-th, but got %v elements", x[i-k+1], i-k+1, freq)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `8 3
NO NO YES YES YES NO
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `9 8
YES NO
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `3 2
NO NO
`
	runSample(t, s)
}
