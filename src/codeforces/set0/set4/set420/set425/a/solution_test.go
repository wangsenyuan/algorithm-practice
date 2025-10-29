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
	s := `10 2
10 -1 2 2 2 2 2 2 -1 10
`
	expect := 32
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 10
-1 -1 -1 -1 -1
`
	expect := -1
	runSample(t, s, expect)
}
