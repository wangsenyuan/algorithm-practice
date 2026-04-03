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
	s := `3
1 1 1`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
4 1 2 1`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
5 10 15 10`
	expect := 15
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1
42`
	expect := 42
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `9
1 2 3 2 2 2 3 2 1`
	expect := 12
	runSample(t, s, expect)
}
