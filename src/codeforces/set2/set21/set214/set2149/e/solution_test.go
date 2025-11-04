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
	s := `1 1 1 1
5`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 2 2 3
1 2 1 3 2`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 3 1 6
1 2 3 1 2 3`
	expect := 10
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 1 1 2
7 7 7 7`
	expect := 7
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `7 3 2 4
1 2 1 2 3 2 1`
	expect := 5
	runSample(t, s, expect)
}
