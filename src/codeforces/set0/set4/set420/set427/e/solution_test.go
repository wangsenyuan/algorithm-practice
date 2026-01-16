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
	s := `3 6
1 2 3`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 5
-7 -6 -3 -1 1
`
	expect := 16
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 369
0
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `11 2
-375 -108 1336 1453 1598 1892 2804 3732 4291 4588 4822
`
	expect := 18716
	runSample(t, s, expect)
}
