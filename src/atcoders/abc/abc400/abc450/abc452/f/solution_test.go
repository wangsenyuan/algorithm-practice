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
	runSample(t, `7 3
6 3 2 1 7 5 4
`, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 1
1 2 3 4
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `25 18
14 19 24 8 12 11 6 5 3 13 22 15 17 2 9 4 7 18 10 25 23 16 1 20 21
`, 3)
}

func TestSample4(t *testing.T) {
	runSample(t, `4 0
1 2 3 4
`, 10)
}
