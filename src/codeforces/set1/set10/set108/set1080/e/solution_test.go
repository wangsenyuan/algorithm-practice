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
	s := `1 3
aba
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 3
aca
aac
`
	expect := 11
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 5
accac
aaaba
cccaa
`
	expect := 43
	runSample(t, s, expect)
}
