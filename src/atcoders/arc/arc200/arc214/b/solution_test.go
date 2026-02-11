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
	s := `2 1
1 2 3`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 0`
	expect := -1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 4
2 1 4
1 3 5
1 4 7
2 3 1`
	expect := 2
	runSample(t, s, expect)
}

