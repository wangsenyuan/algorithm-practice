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
	s := `4
2 4
4 1
3 1
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10
7 1
8 4
8 10
4 7
6 5
9 3
3 5
2 10
2 5
`
	expect := 4
	runSample(t, s, expect)
}


func TestSample3(t *testing.T) {
	s := `2
1 2
`
	expect := 0
	runSample(t, s, expect)
}
