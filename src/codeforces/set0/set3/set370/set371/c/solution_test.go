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
	s := `BBBSSC
6 4 1
1 2 3
4
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `BBC
1 10 1
1 10 1
21
`
	expect := 7
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `BSC
1 1 1
1 1 3
1000000000000
`
	expect := 200000000001

	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `B
1 1 1
1 1 1
381
`
	expect := 382

	runSample(t, s, expect)
}
