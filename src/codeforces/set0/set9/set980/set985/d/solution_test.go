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
	s := "5 2"
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "6 8"
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "20 4"
	expect := 7
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "1000000000000000000 1000000000000000000"
	expect := 1414213562
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "1000000000000000000 1"
	expect := 1999999999
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "30 4"
	expect := 8
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := "696616491401388220 958775125"
	expect := 1191798158
	runSample(t, s, expect)
}
