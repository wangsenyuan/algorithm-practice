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
	runSample(t, `7
4 1
4 2
4 2
2 1
5 4
6 4
3 2
`, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, `10
17 15
5 4
18 16
13 12
19 17
3 1
12 10
5 3
18 16
11 10
`, 1)
}
