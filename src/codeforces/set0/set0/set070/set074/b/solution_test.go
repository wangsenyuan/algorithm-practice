package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 3 2
to head
0001001
`
	expect := "Stowaway"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 2 1
to tail
0001
`
	expect := "Controller 2"
	runSample(t, s, expect)
}
