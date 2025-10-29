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
	s := `3 2
10 8 1
2 7 1
`
	expect := 18
	runSample(t, s, expect)
}


func TestSample2(t *testing.T) {
	s := `5 3
4 4 4 4 4
2 2 2 2 2
`
	expect := -1
	runSample(t, s, expect)
}
