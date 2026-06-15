package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	X, Y, a := drive(reader)

	if len(a) != 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, a)
	}

	if !expect {
		return
	}

	ok := make([]bool, len(Y))
	for i := range a {
		var row bool
		for j, v := range a[i] {
			if v > X[i] || v > Y[j] {
				t.Fatalf("Sample result %v, is not valid", a)
			}
			if v == Y[j] {
				ok[j] = true
			}
			if v == X[i] {
				row = true
			}
		}
		if !row {
			t.Fatalf("Sample result %v is invalid, it has no %d in row %d", a, X[i], i)
		}
	}

	for j := range Y {
		if !ok[j] {
			t.Fatalf("Sample result %v is invalid, it has no %d in col %d", a, Y[j], j)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `2 3
5 6
5 3 6`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
5 4 6
6 2 4`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 4
18 20 19 14 17
18 20 14 15`
	expect := true
	runSample(t, s, expect)
}
