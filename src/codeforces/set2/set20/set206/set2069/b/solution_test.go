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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1 1
1`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
1 2 1
2 3 2
1 3 1`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 6
5 4 5 4 4 5`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 4
1 4 2 2
1 4 3 5
6 6 3 5`
	expect := 10
	runSample(t, s, expect)
}
