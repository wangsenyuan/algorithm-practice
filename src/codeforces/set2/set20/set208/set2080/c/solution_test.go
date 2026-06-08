package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 1
5 3
1 2
4
`
	runSample(t, s, "First")
}

func TestSample2(t *testing.T) {
	s := `1 2
2
3
4 1
`
	runSample(t, s, "Second")
}
