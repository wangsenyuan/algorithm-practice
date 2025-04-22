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
	s := `1 1
0`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 1
0 1 3 2 4`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 2
2 1 0 0 1 2`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 5
0 0 0 0 0`
	expect := 1
	runSample(t, s, expect)
}