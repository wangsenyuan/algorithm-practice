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
	runSample(t, `5
1 2 3 2 1
`, 9)
}

func TestSample2(t *testing.T) {
	runSample(t, `7
5 4 1 1 1 1 3
`, 37)
}

func TestSample3(t *testing.T) {
	runSample(t, `6
1 2 3 4 5 6
`, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, `6
4 1 6 3 2 6
`, 17)
}

func TestSample5(t *testing.T) {
	runSample(t, `7
1 3 2 7 2 3 1
`, 29)
}
