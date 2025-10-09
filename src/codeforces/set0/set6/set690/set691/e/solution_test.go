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
15 1 2 4 8`
	expect := 13
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 1
15 1 2 4 8`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 2
93 93 85 48 44 98 93 100 98 98`
	expect := 52
	runSample(t, s, expect)
}
