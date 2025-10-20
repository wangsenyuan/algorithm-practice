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
	s := `6
3 1 4 -1 -5 -9`
	expect := 23
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
-10 -3 -17 1 19 20`
	expect := 40
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `15
-19 -16 -6 5 -9 -14 16 18 16 6 6 -14 -18 -8 13`
	expect := 107
	runSample(t, s, expect)
}
