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
	s := `3
4 4 8
1 5 0
5 2 10
	`
	expect := 9
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
6 2 4
1 6 2
2 4 3
5 3 8
	`
	expect := 10
	runSample(t, s, expect)
}
