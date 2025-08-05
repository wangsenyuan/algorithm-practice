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
	s := `1 3 3`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 3 10`
	expect := 165
	runSample(t, s, expect)
}
