package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4 1
XX#X
#XX#
#X#X`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 4 2
XX#X
#XX#
#X#X`
	expect := 60
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 1 3
X
X
#`
	expect := 0
	runSample(t, s, expect)
}
