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
	s := `2 3
1
1`
	expect := 18
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 4
3
1 4 6`
	expect := 33664
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `12 31
6
1 3 5 7 9 11`
	expect := 909076242
	runSample(t, s, expect)
}
