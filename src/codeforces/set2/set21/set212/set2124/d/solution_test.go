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
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 3
5 4 3 4 5`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 1
1 1 2 1`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 6
2 3 4 5 3 2`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 4
5 2 4 3 1`
	expect := false
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `8 5
4 7 1 2 3 1 3 4`
	expect := false
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `5 4
1 2 1 2 2`
	expect := true
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `3 3
1 2 2`
	expect := false
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `4 4
2 1 2 2`
	expect := true
	runSample(t, s, expect)
}
