package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	k, a, res := drive(reader)

	if res[1]-res[0]+1 != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}

	freq := make(map[int]int)
	for i := res[0] - 1; i < res[1]; i++ {
		freq[a[i]]++
	}

	if len(freq) > k {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 5
1 2 3 4 5
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `9 3
6 5 1 2 3 2 1 4 5
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 1
1 2 3
`
	expect := 1
	runSample(t, s, expect)
}
