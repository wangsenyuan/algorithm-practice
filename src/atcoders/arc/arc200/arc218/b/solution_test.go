package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1
2`
	expect := "Alice"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 1 1`
	expect := "Bob"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
1 2 3 4`
	expect := "Bob"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `7
3 1 4 1 5 9 2`
	expect := "Bob"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3
218 503 2026`
	expect := "Alice"
	runSample(t, s, expect)
}
