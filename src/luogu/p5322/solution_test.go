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
	s := `1 3 10
2 2 6`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 3 10
2 2 6
0 0 0`
	expect := 8
	runSample(t, s, expect)
}
