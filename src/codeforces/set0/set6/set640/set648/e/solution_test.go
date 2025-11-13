package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 3
123 1
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 10
1
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 4
1 2 3
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 777
12 23 345
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `10 137
36 3 4078 482192 4 60354 562127960 32271816 1700 2612
`
	expect := 7
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `1 2
1
`
	expect := 0
	runSample(t, s, expect)
}
