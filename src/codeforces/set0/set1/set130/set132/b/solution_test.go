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
	s := `2 10
12
43
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 12
1423
6624
6625
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 9
10345
23456
34567
45678
56789
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 9995
11122
06330
04470
55800
`
	expect := 3
	runSample(t, s, expect)
}
