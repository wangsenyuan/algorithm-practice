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
	s := `3 2
30000 15000
40000 15000
50000 15000
`
	expect := 49999
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 1
1 20000
10000 10000
20000 20000
25000 10000
80000 60000
`
	expect := 39999
	runSample(t, s, expect)
}
