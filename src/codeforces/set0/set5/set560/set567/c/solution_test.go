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
	s := `5 2
1 1 2 2 4`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1
1 1 1`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 3
1 2 6 2 3 6 9 18 3 9`
	expect := 6
	runSample(t, s, expect)
}
