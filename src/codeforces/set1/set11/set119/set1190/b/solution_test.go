package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1
0
`
	expect := "cslnb"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
1 0
`
	expect := "cslnb"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
2 2
`
	expect := "sjfnb"
	runSample(t, s, expect)
}


func TestSample4(t *testing.T) {
	s := `3
4 4 4
`
	expect := "cslnb"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3
0 0 6
`
	expect := "cslnb"
	runSample(t, s, expect)
}