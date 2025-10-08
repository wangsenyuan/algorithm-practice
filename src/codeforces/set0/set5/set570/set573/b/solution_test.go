package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := drive(reader)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `6
2 1 4 6 2 2
`
	runSample(t, s, 3)
}

func TestSample2(t *testing.T) {
	s := `7
3 3 3 1 3 3 3
`
	runSample(t, s, 2)
}
