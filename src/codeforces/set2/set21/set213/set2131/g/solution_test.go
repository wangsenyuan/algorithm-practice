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
	s := `2 3
1 3`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 6
5 1 4`
	expect := 24
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 100
2 100`
	expect := 118143737
	runSample(t, s, expect)
}
