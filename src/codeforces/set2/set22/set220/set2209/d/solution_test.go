package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expectLength int) {
	t.Helper()

	reader := bufio.NewReader(strings.NewReader(input))
	var r, g, b int
	fmt.Fscan(reader, &r, &g, &b)
	reader = bufio.NewReader(strings.NewReader(input))
	res := drive(reader)

	if len(res) != expectLength {
		t.Fatalf("expected length %d, but got %d (%q)", expectLength, len(res), res)
	}

	used := map[byte]int{'R': 0, 'G': 0, 'B': 0}
	limits := map[byte]int{'R': r, 'G': g, 'B': b}
	for i := range len(res) {
		if _, ok := used[res[i]]; !ok {
			t.Fatalf("invalid color %q in %q", res[i], res)
		}
		used[res[i]]++
		if used[res[i]] > limits[res[i]] {
			t.Fatalf("uses too many %c ghostfires in %q", res[i], res)
		}
		if i > 0 && res[i] == res[i-1] {
			t.Fatalf("adjacent equal ghostfires in %q", res)
		}
		if i >= 3 && res[i] == res[i-3] {
			t.Fatalf("equal ghostfires at distance three in %q", res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "0 0 1\n", 1)
}

func TestSample2(t *testing.T) {
	runSample(t, "1 1 1\n", 3)
}

func TestSample3(t *testing.T) {
	runSample(t, "0 3 0\n", 1)
}

func TestSample4(t *testing.T) {
	runSample(t, "2 2 2\n", 6)
}

func TestSample5(t *testing.T) {
	runSample(t, "2 7 3\n", 11)
}

func TestPrefersTwoBackColorOnTie(t *testing.T) {
	runSample(t, "1 1 2\n", 4)
}
