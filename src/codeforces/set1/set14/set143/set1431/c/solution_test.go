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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 2
1 3 4 6 7`
	expect := 7
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 3
1 2 3 4 5 6`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 3
3 3 4 4 5 5`
	expect := 6
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1 1
7`
	expect := 7
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4 4
1 3 3 7`
	expect := 1
	runSample(t, s, expect)
}
