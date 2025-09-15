package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	total, _ := drive(reader)

	if expect != total {
		t.Fatalf("Sample expect %d, but got %d", expect, total)
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
1 2 1
2 3 1
1 3 2
3
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 4
1 2 1
2 3 1
3 4 1
4 1 2
4
`
	expect := 4
	runSample(t, s, expect)
}
