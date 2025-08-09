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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 10
10
21
27
27
`
	runSample(t, s, 3)
}

func TestSample2(t *testing.T) {
	s := `8 8
9
19
28
36
45
55
66
78
`
	runSample(t, s, 6)
}

func TestSample3(t *testing.T) {
	s := `13 7
8
8
9
16
17
17
18
21
23
24
24
26
30
`
	runSample(t, s, 4)
}
