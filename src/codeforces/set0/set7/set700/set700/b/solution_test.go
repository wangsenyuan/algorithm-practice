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
	s := `7 2
1 5 6 2
1 3
3 2
4 5
3 7
4 3
4 6
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `9 3
3 2 1 6 5 9
8 9
3 2
2 7
3 4
7 6
4 5
2 1
2 8
`
	expect := 9
	runSample(t, s, expect)
}
