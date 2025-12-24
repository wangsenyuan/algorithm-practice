package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	A, c, B := drive(reader)

	if B == expect {
		return
	}

	if B == -1 || expect == -1 {
		t.Fatalf("Sample expect %d, but got %d", expect, B)
	}
	freq := make([]int, 2)
	for _, v := range c {
		if v == A {
			freq[0]++
		} else if v == B {
			freq[1]++
		}
		if freq[0] > freq[1] {
			t.Fatalf("Sample expect %d, but got %d", expect, B)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4 1
2 1 4 2
`
	expect := 2
	runSample(t, s, expect)
}


func TestSample2(t *testing.T) {
	s := `5 2
2 2 4 5 3
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 10
1 2 3
`
	expect := 4
	runSample(t, s, expect)
}