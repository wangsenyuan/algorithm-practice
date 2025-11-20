package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
1 2
1 2
1 0
0 1
0 2
`
	expect := false
	runSample(t, s, expect)
}


func TestSample2(t *testing.T) {
	s := `2
1 3
1 0
-1 3
`
	expect := true
	runSample(t, s, expect)
}


func TestSample3(t *testing.T) {
	s := `2
1 3
1 0
0 2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3
0 2
10 0
0 0
8 2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3
4 11
-2 14
2 -15
-8 -15
`
	expect := true
	runSample(t, s, expect)
}
