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
	s := `1
5*a++-3*++a+a++
`
	expect := 11
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
a+++++a
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `-668
820*a+++402*++a-482*++a
`
	expect := -492358
	runSample(t, s, expect)
}
