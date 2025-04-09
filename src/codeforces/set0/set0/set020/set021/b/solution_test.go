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
	s := `1 1 0
2 2 0
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 1 0
2 -2 0
`
	expect := 1
	runSample(t, s, expect)
}
func TestSample3(t *testing.T) {
	s := `-1 -1 0
1 1 -1
`
	expect := 0
	runSample(t, s, expect)
}
