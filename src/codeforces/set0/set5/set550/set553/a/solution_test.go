package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
2
2
1
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
1
2
3
4
`
	expect := 1680
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
100
100
100
100
100
100
100
100
100
100
`
	expect := 12520708
	runSample(t, s, expect)
}