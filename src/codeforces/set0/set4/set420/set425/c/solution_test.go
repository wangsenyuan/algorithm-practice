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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 5 100000 1000
1 2 3 4 5
3 2 4 5 1`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 4 3006 1000
1 2 3
1 2 4 3`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 19 100000 1796
5
3 3 3 5 3 2 4 2 5 5 2 2 4 3 2 4 2 5 2`
	expect := 1
	runSample(t, s, expect)
}
