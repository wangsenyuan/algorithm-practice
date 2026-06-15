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
	s := `5
1 2 3 4 5
5 4 3 2 0`
	expect := 25
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
1 2 3 10 20 30
6 5 4 3 2 0`
	expect := 138
	runSample(t, s, expect)
}
