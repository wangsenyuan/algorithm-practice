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
	s := `3 4 3
2 1 1 1
1 1 1 1
2 1 1 3
	`
	runSample(t, s, 5)
}

func TestSample2(t *testing.T) {
	s := `3 3 9
1 3 5
8 9 7
4 6 2
	`
	runSample(t, s, 22)
}

func TestSample3(t *testing.T) {
	s := `3 4 12
1 2 3 4
8 7 6 5
9 10 11 12
	`
	runSample(t, s, 11)
}

func TestSample4(t *testing.T) {
	s := `5 5 5
4 2 1 2 3
3 4 4 2 2
3 4 1 2 4
2 1 5 4 2
4 3 1 1 2
	`
	runSample(t, s, 9)
}
