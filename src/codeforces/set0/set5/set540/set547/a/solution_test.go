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
4 2
1 1
0 1
2 3`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {

	s := `1023
1 2
1 0
1 2
1 1`
	expect := -1
	runSample(t, s, expect)
}
