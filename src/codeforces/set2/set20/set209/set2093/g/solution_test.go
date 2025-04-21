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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 0
1 2 3 4 5`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 7
1 2 3 4 5`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 8
1 2 3 4 5`
	expect := -1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 7
3 5 1 4 2`
	expect := 4
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5 3
3 5 1 4 2`
	expect := 2
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `6 71
26 56 12 45 60 27`
	expect := -1
	runSample(t, s, expect)
}
