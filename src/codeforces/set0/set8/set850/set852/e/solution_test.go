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
		t.Errorf("expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
1 2`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 2
2 3`
	expect := 10
	runSample(t, s, expect)
}
