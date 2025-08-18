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
	s := `5
3 2 1 6 5`
	expect := 8
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
3 3 3 3`
	expect := 0
	runSample(t, s, expect)
}
