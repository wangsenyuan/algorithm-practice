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
	s := `5
1 4 3 2 5`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 2 2 2 1`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7
10 20 40 50 70 90 30`
	expect := 0
	runSample(t, s, expect)
}
