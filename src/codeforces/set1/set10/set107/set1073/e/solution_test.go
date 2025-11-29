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
	s := `10 50 2`
	expect := 1230
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 50 1`
	expect := 110
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 2345 10`
	expect := 2750685
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `101 154 2`
	expect := 2189
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `427896435961371452 630581697708338740 1`
	expect := 716070897
	runSample(t, s, expect)
}
