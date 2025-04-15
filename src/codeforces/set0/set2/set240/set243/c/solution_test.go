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
	s := `5
R 8
U 9
L 9
D 8
L 2
`
	expect := 101
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7
R 10
D 2
L 7
U 9
D 2
R 3
D 10
`
	expect := 52
	runSample(t, s, expect)
}
