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
	s := `3 7
1 1
2 2
3 1
1 3
3 2
2 3
3 3
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 6
1 1
2 2
3 1
3 2
2 3
3 3
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 1
2 2
`
	expect := 6
	runSample(t, s, expect)
}
