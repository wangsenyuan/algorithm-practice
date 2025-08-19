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
	s := `3 3
UUU
L?R
DDD`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 3
???
???`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 3
?U?
R?L
RDL`
	expect := 5
	runSample(t, s, expect)
}
