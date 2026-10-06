package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, possible bool) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	var n, k int
	if _, err := fmt.Fscan(reader, &n, &k); err != nil {
		t.Fatal(err)
	}
	a := make([]int, n)
	for i := range a {
		if _, err := fmt.Fscan(reader, &a[i]); err != nil {
			t.Fatal(err)
		}
	}

	ends := drive(bufio.NewReader(strings.NewReader(input)))
	if !possible {
		if ends != nil {
			t.Fatalf("expected NO, got %v", ends)
		}
		return
	}
	if len(ends) != k {
		t.Fatalf("expected %d segment ends, got %v", k, ends)
	}
	start := 0
	for _, end := range ends {
		if end <= start || end > n {
			t.Fatalf("invalid segment ends: %v", ends)
		}
		sum := 0
		for _, x := range a[start:end] {
			sum += x
		}
		if sum%2 == 0 {
			t.Fatalf("segment %d..%d has even sum", start+1, end)
		}
		start = end
	}
	if start != n {
		t.Fatalf("segments end at %d, want %d", start, n)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "5 3\n7 18 3 14 1\n", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "5 4\n1 2 3 4 5\n", false)
}

func TestSample3(t *testing.T) {
	runSample(t, "6 2\n1 2 8 4 10 2\n", false)
}
