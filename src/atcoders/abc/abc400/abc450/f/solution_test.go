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
	s := `4 3
1 4
1 3
2 4
`
	expect := 5

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 11
1 4
1 4
3 9
2 5
3 4
9 10
6 9
4 10
1 3
8 10
4 7
`
	expect := 1297
	runSample(t, s, expect)
}
