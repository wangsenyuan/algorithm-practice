package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 5 4 6 1 3 6 2 5 5 1 2 3 5 3 1 1 2 4 6 6 4 3 4`
	expect := false
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 3 5 3 2 5 2 5 6 2 6 2 4 4 4 4 1 1 1 1 6 3 6 3`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 2 1 1 5 5 5 5 3 3 4 4 1 4 1 4 2 3 2 3 6 6 6 6`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1 1 1 1 3 3 3 3 5 5 5 5 2 2 2 2 4 4 4 4 6 6 6 6`
	expect := false
	runSample(t, s, expect)
}
