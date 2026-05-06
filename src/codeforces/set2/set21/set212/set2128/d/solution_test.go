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
3 2 1
	`
	runSample(t, s, 10)
}

func TestSample2(t *testing.T) {
	s := `4
4 3 1 2
	`
	runSample(t, s, 17)
}

func TestSample3(t *testing.T) {
	s := `6
6 1 5 2 4 3
	`
	runSample(t, s, 40)
}

func TestSample4(t *testing.T) {
	s := `3
2 3 1
	`
	runSample(t, s, 8)
}
