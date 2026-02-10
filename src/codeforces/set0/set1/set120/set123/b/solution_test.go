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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 2 1 0 0 1`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2 10 11 0 1`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 4 3 -1 3 7`
	expect := 2
	runSample(t, s, expect)
}
