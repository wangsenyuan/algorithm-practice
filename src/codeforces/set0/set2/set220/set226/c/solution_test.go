package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `10 1 8 2`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 1 8 3`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 1 1000000000000 1000000000000`
	expect := 0
	runSample(t, s, expect)
}
