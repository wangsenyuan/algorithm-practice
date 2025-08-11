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
	s := `7
3 1 6 7
4 3 5 2 9
2 8 1
4 3 7 6 4
3 2 5 9
3 6 3 8
3 4 2 9
`
	expect := 6
	runSample(t, s, expect)
}
