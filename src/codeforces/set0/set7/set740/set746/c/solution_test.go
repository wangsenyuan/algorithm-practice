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
	s := `4 2 4
3 4
1 1
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 4 0
1 2
3 1
`
	expect := 7
	runSample(t, s, expect)
}

func TestTramAlreadyAtX1MovingLeft(t *testing.T) {
	s := `10 5 3
1 2
5 -1
`
	expect := 2
	runSample(t, s, expect)
}


func TestSample3(t *testing.T) {
	s := `5 4 0
5 14
1 -1
`
	expect := 55
	runSample(t, s, expect)
}
