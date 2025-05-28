package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readNum(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 3 1000000
6`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 3 2
4`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `9 8 7
8`
	runSample(t, s)
}
