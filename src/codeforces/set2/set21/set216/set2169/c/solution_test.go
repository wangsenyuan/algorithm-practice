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
	s := `3
2 5 1
	`
	runSample(t, s, 13)
}

func TestSample2(t *testing.T) {
	s := `2
4 4
	`
	runSample(t, s, 8)
}

func TestSample3(t *testing.T) {
	s := `4
1 3 2 1
	`
	runSample(t, s, 20)
}

func TestSample4(t *testing.T) {
	s := `5
3 2 0 9 10
	`
	runSample(t, s, 32)
}
