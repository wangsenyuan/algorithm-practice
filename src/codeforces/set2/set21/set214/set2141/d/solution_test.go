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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 16
1 10 2
`, -1)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 20
6 2 4 9
`, 11)
}

func TestSample3(t *testing.T) {
	runSample(t, `5 9
7 7 7 7 7
`, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, `2 1000000000000
1000000000 1000000000
`, 499999999999)
}

