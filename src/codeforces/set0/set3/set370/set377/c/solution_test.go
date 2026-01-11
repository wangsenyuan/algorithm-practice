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
	s := `2
2 1
2
p 1
p 2
`
	runSample(t, s, 1)
}

func TestSample2(t *testing.T) {
	s := `6
6 4 5 4 5 5
4
b 2
p 1
b 1
p 2
`
	runSample(t, s, 0)
}

func TestSample3(t *testing.T) {
	s := `4
1 2 3 4
4
p 2
b 2
p 1
b 1
`
	runSample(t, s, -2)
}
