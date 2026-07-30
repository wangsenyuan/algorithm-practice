package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
1
2
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
2 4 5
1 3 6
`, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, `4
7 5 4 8
4 6 7 8
`, 6)
}

func TestSample4(t *testing.T) {
	runSample(t, `8
8 7 13 11 1 10 4 5
11 11 12 8 9 2 3 13
`, 8)
}

func TestSample5(t *testing.T) {
	runSample(t, `9
16 1 9 12 5 18 10 10 16
14 6 7 11 12 17 18 3 17
`, 14)
}

func TestSample6(t *testing.T) {
	runSample(t, `6
3 6 12 4 10 12
2 3 2 7 8 9
`, 8)
}
