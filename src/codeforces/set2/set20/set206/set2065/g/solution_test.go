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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
2 2 3 4`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
2 2 3 4 5 6`
	expect := 12
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `9
2 2 4 5 7 8 9 3 5`
	expect := 18
	runSample(t, s, expect)
}
