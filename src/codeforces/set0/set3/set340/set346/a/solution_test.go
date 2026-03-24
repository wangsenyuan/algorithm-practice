package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
2 3
`
	expect := "Alice"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
5 3
`
	expect := "Alice"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
5 6 7
`
	expect := "Bob"
	runSample(t, s, expect)
}