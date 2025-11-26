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
	runSample(t, `3 2
50 4 20
`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 3
15 16 3 25 9
`, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 3
9 77 13
`, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, `1 1
500
`, 2)
}

func TestSample5(t *testing.T) {
	runSample(t, `2 2
2199023255552 11920928955078125
`, 23)
}
