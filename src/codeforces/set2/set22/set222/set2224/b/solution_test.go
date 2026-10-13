package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
0 0 0 0 0
`, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
0 1
`, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
1 1 1 1 0
`, 13)
}

func TestSample4(t *testing.T) {
	runSample(t, `6
1 1 4 5 1 4
`, 30)
}

func TestSample5(t *testing.T) {
	runSample(t, `1
1000000000
`, 1000000000)
}

func TestSample6(t *testing.T) {
	runSample(t, `20
2 12 3 5 14 0 7 11 0 11 4 6 14 8 2 13 9 1 10 2
`, 461)
}
