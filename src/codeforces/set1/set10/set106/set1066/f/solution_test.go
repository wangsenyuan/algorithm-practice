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
	s := `8
2 2
1 4
2 3
3 1
3 4
1 1
4 3
1 2
`
	expect := 15
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
2 1
1 0
2 0
3 2
0 3
`
	expect := 9
	runSample(t, s, expect)
}
