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
	s := `2 2 1`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2 3`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 2 2`
	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 2 4`
	expect := 0
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `1 1 1`
	expect := 0
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `3 1 1`
	expect := 1
	runSample(t, s, expect)
}