package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `4 3
2 1 4 3
`
	// 2 1 4 3 2 1 4 3
	// 2 3 7 10 12 13 17 20
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 1000000
1 1
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `9 5
9 9 8 2 4 4 3 5 3
`
	// 8 + 2
	// 2 + 4 + 4
	// 5
	// 4 + 3 + 5 + 3
	// 9 + 9 + 8 + 2 + 4 + 4 = 30
	expect := 11
	runSample(t, s, expect)
}
