package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int) {
	t.Helper()

	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
1 3 4 2
`, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, `1
1
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `20
11 10 18 13 12 16 5 19 7 6 17 4 9 1 14 2 20 15 8 3
`, 431610)
}
