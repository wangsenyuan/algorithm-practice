package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
7 5 8 6 9`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
-1 -2 0 0 -5`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
5 4 3 2 1`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
1000000000 0 0 0 0`
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5
-1 -1 -1 -1 1`
	expect := true
	runSample(t, s, expect)
}
