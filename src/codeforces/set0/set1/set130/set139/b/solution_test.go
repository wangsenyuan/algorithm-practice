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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1
5 5 3
3
10 1 100
15 2 320
3 19 500
`
	expect := 640
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
29 30 29
30 15 28
27 30 23
3
30 27 21
11 24 30
25 20 12
`
	expect := 261
	runSample(t, s, expect)
}
