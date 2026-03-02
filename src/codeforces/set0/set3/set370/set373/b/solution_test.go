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
	s := `9 1 1`
	expect := 9
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `77 7 7`
	expect := 7
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `114 5 14`
	expect := 6
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1 1 2`
	expect := 0
	runSample(t, s, expect)
}