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
	s := `5
1 1 3 3
1 2 1 2 3
1 1 3 1 1
`
	expect := 144
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
1
1 1
2 1
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
3 1
1000000000 1 1
1 1 1
`
	expect := 1755647
	runSample(t, s, expect)
}

func TestLargeAvailableCandiesDoNotOverflow(t *testing.T) {
	s := `10
1 1 1 1 1 1 1 1 1
1000000000 1000000000 1000000000 1000000000 1000000000 1000000000 1000000000 1000000000 1000000000 1000000000
6 1 1 1 1 1 1 1 1 1
`
	expect := 553107633
	runSample(t, s, expect)
}
