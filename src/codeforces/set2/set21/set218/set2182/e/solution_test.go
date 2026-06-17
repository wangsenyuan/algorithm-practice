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
	s := `2 1 6
1
1 2 3
1 2 7`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2 3
1 1
2 1 3
2 1 5`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 4 11
1 2 2 1
3 2 5
4 4 6
3 1 3`
	expect := 2
	runSample(t, s, expect)
}
