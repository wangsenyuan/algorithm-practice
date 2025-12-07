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
	s := `2
5 23 1 2
3 13 2 3
	`
	runSample(t, s, 2)
}

func TestSample2(t *testing.T) {
	s := `3
12 9 2 1
12 8 1 3
12 8 2 2
	`
	runSample(t, s, 3)
}

func TestSample3(t *testing.T) {
	s := `1
1 10 1 13
	`
	runSample(t, s, 1)
}