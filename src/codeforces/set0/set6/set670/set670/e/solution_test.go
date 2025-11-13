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
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `8 4 5
(())()()
RDLD
`
	expect := "()"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `12 5 3
((()())(()))
RRDLD
`
	expect := "(()(()))"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8 8 8
(())()()
LLLLLLDD
`
	expect := "()()"
	runSample(t, s, expect)
}
