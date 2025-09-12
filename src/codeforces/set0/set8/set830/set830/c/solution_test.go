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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4
1 3 5`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 40
10 30 50`
	expect := 32
	runSample(t, s, expect)
}
