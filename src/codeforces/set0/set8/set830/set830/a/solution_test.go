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
	s := `2 4 50
20 100
60 10 40 80
`
	expect := 50
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 2 10
11
15 7
`
	expect := 7
	runSample(t, s, expect)
}
