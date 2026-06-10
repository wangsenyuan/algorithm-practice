package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))

	b, a := drive(reader)

	if len(a) != len(b)+1 {
		t.Fatalf("Sample result %v, is invalid", a)
	}

	freq := make(map[int]int)
	var diff int
	for i, v := range a {
		if i&1 == 0 {
			diff += v
		} else {
			diff -= v
		}
		freq[v]++
		if freq[v] > 1 {
			t.Fatalf("Sample result %v, is invalid, it has duplicates %d", a, v)
		}
	}

	if diff != 0 {
		t.Fatalf("Sample result %v, sum(odd) != sum(even)", a)
	}

	for i, v := range b {
		freq[v]--
		if freq[v] < 0 {
			t.Fatalf("Sample result %v is invalid, it doesn't contain all b[%d] %d", a, i, v)
		}
	}

}

func TestSample1(t *testing.T) {
	s := `1
9 2`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `2
8 6 1 4`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `3
99 2 86 33 14 77`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `2
1 6 3 2`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `1
1 2`
	runSample(t, s)
}

