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
	s := `6
1 6 2 5 3 7`
	expect := 12
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
6 6 6 0 0 0`
	expect := 27
	runSample(t, s, expect)
}
