package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "2 1 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0"
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "3 1 1 0 0 0 0 0 0 0 0 0 0 0 0 1 0 1 0 0 0 0 0 0 1 0"
	expect := 960
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "0 0 0 0 0 0 0 0 0 0 0 0 1 0 3 0 0 0 0 0 0 0 0 0 0 0"
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 233527 233827"
	expect := 789493841
	runSample(t, s, expect)
}