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
	s := `10 5 6
3 2 0 5 6 1
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 6 4
3 5 0 4
`	
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `257 0 3
0 0 256
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 0 20
1 1 1 0 1 1 1 1 0 0 0 0 1 0 0 0 0 1 0 1
`
	expect := 22
	runSample(t, s, expect)
}
