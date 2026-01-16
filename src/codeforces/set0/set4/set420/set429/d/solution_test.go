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
	s := `4
1 0 0 -1
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "4\n1 2 3 4"
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
1 -1
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4
200 100 -200 100
`
	expect := 9
	runSample(t, s, expect)
}
