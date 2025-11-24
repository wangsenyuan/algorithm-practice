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
	s := `4 6
6 1 2 4
3 5 1 7 2 3`
	expect := 16
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 5
9 10 6 3 7
5 9 2 3 9`
	expect := 25
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 6
5 8
8 7 5 8 2 10`
	expect := 11
	runSample(t, s, expect)
}
