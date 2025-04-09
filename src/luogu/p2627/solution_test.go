package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 2
1
2
3
4
5
`
	expect := 12
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 2
1
2 
5 
6 
2
`
	expect := 12
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 2
7
2 
1
1
100 
8
`
	expect := 117
	runSample(t, s, expect)
}
