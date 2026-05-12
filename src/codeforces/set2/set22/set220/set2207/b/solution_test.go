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
	s := `1 2 10
10
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 1 32
1 4 9 16 25
`
	expect := 7
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 3 40
13 37
`
	expect := 19
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 2 7
6 7
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `8 5 60
3 17 20 28 36 44 45 50
`
	expect := 19
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `6 7 1987
6 7 66 77 666 777
`
	expect := 1477
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `1 1 1
1
`
	expect := 0
	runSample(t, s, expect)
}
