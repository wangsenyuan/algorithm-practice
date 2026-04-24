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
	s := `7 4
1 3 1 2 3 1 1
`
	expect := 14
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8 10
7 5 1 3 2 5 6 8
`
	expect := 43
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 1
1
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 20
5 2 1 2 1 3 6 7 1 1
`
	expect := 170
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5 1000000000
1 420420420 1 420420420 1
`
	expect := 3738738738
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `1 1000000000
1
`
	expect := 999999999
	runSample(t, s, expect)
}

