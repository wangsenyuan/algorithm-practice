package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `3
5 4 5`
	expect := 12
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
4 5 4`
	expect := 11
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
3 3 3 3 4 1 2 3 5 4`
	expect := 12
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10
17 89 92 42 29 41 92 14 70 45`
	expect := 186
	runSample(t, s, expect)
}
