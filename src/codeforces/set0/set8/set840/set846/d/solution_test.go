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
	s := `2 3 2 5
2 1 8
2 2 8
1 2 1
1 3 4
2 3 2
`
	// -1 1 4
	// 8 8 2
	expect := 8
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3 2 5
1 2 2
2 2 1
2 3 5
3 2 10
2 1 100
`
	expect := -1
	runSample(t, s, expect)
}
