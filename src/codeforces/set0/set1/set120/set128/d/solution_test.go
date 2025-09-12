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
	s := `4
1 2 3 2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
1 1 2 2 2 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
2 4 1 1 2 2
`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
6 7 6 7 6
`
	expect := false
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `8
3 5 8 4 7 6 4 7
`
// 6787
	expect := false
	runSample(t, s, expect)
}
