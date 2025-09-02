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
	s := `3 1 2 3`
	// 考虑 v = 4 时， (1, 3), (2, 4)
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 3 2 3`
	expect := 4
	runSample(t, s, expect)
}
