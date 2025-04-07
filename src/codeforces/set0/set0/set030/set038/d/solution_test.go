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
	s := `2
0 0 3 3
1 0 4 3
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
0 0 3 3
2 0 5 3
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
0 0 3 3
1 0 4 3
2 0 5 3
`
	expect := 3
	runSample(t, s, expect)
}


func TestSample4(t *testing.T) {
	s := `5
7 -10 -8 5
4 -7 -5 2
2 -5 -3 0
-9 48 50 -11
50 -4 -2 48
`
	expect := 3
	runSample(t, s, expect)
}