package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	_, res := process(reader)
	if len(res) > 0 != (expect != "-1") {
		t.Fatalf("Sample expect %s, but got %v", expect, res)
	}
	if expect == "-1" {
		return
	}
	s = fmt.Sprintf("%v", res)
	s = s[1 : len(s)-1]
	if s != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, s)
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
1 2 1
`
	expect := "2"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 10
3 3 2 1
`
	expect := "-1"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7 10
1 3 3 1 2 3 1
`
	expect := "6 2 3"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1 0
1
`
	expect := "1"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `6 101
9 78 54 62 2 91
`
	expect := "4 6 2 3"
	runSample(t, s, expect)
}
