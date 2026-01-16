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
	s := `4
17 -14
52 -5
1 52
6 0
	`
	runSample(t, s, 4)
}

func TestSample2(t *testing.T) {
	s := `5
4 5
3 2
5 -3
6 -2
4 3
	`
	runSample(t, s, 14)
}

func TestSample3(t *testing.T) {
	s := `5
4 -5
5 -2
3 -1
4 -1
5 -4
	`
	runSample(t, s, 14)
}
