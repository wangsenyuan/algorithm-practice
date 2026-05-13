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
	s := `10
10 9 8 7 6 5 4 3 2 1
`
	expect := "Alice"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 8192 677
`
	expect := "Bob"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
6 7
`
	expect := "Bob"
	runSample(t, s, expect)
}
