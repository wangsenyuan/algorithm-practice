package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
67 67`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 2 3`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
67 667 167 867 267 467 367 567 767 967`
	expect := 1012
	runSample(t, s, expect)
}
