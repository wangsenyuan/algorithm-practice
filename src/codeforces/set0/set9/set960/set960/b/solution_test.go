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
	s := `2 0 0
1 2
2 3`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 1 0
1 2
2 2`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 5 7
3 4
14 4`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 5 5
0 0 0 0 0
0 0 0 0 0`
	expect := 0
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `2 1 1
0 0
1 1`
	expect := 0
	runSample(t, s, expect)
}
