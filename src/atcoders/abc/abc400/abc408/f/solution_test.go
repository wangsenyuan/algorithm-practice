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
	s := `5 2 1
5 3 1 4 2
`
	expect := 2
	runSample(t, s, expect)
}


func TestSample2(t *testing.T) {
	s := `13 3 2
13 7 10 1 9 5 4 11 12 2 8 6 3
`
	expect := 3
	runSample(t, s, expect)
}
