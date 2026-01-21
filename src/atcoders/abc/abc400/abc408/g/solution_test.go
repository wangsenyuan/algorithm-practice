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
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 2 2 1`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 2 8 3`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 2 2 1`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `60 191 11 35`
	expect := 226
	runSample(t, s, expect)
}
