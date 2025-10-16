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
	s := `3 2 2
2 4 1`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 4 2
10 1 6 2 5`
	expect := 18
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 2 3
3 1 4 2 5`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 4 4
9 5 2 2`
	expect := 18
	runSample(t, s, expect)
}
