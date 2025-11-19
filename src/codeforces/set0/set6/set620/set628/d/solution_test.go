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
	s := `2 6
10
99
	`
	runSample(t, s, 8)
}

func TestSample2(t *testing.T) {
	s := `2 0
1
9
	`
	runSample(t, s, 4)
}

func TestSample3(t *testing.T) {
	s := `19 7
1000
9999
	`
	runSample(t, s, 6)
}
