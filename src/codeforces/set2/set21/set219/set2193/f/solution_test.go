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
	s := `1 2 3 5 2
4
4`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1 3 5 2
3 4 3
5 4 1`
	expect := 13
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 1 2 7 3
5 2 3 5 5 3
6 4 3 1 4 1`
	expect := 19
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 6 9 8 6
7 7 7 7 7
3 1 8 8 3`
	expect := 15
	runSample(t, s, expect)
}
