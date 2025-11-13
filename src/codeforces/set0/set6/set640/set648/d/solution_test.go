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
	s := `5 4 
-2 0 4 8 13 
-1 1 
4 3 
6 3 
11 2
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3 
-1 3 7 
1 1 
4 1 
7 1
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 4 
20 1 10 30 
1 1 
2 5 
22 2 
40 10
`
	expect := 3
	runSample(t, s, expect)
}
