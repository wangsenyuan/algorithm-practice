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
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
1 14 5 1 4
1 1 1 1 1
`, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
3 2 5
3 2 5
`, 665496236)
}

func TestSample3(t *testing.T) {
	runSample(t, `10
10 72 65 43 73 23 78 13 49 99
31 90 45 19 44 18 59 31 48 29
`, 820778710)
}
