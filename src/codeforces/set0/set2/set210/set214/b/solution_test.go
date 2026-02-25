package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1
0
`
	expect := "0"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `11
3 4 5 4 5 3 5 3 4 4 0
`
	expect := "5554443330"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8
3 2 5 1 5 2 2 3
`
	expect := "-1"
	runSample(t, s, expect)
}