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
	s := `2
1000 1`
	expect := 999
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
9 9 9 9 9`
	expect := 13
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
7 1 8 4`
	expect := 12
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6
1 14 1 14 1 15`
	expect := -7
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `9
31 12 14 22 89 6 78 25 91`
	expect := 265
	runSample(t, s, expect)
}
