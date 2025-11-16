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
	s := `0 1 0 0 0 0 0`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `0 2 0 0 0 0 0`
	expect := 7
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 1 1 0 0 0 0`
	expect := 9
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 1 0 3 0 0 1`
	expect := 411199181
	runSample(t, s, expect)
}
