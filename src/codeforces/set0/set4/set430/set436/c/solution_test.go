package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	tot, _ := drive(reader)
	if tot != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, tot)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 3 3 2
A.A
...
A.a
..C
X.Y
...
`, 14)
}

func TestSample2(t *testing.T) {
	runSample(t, `1 1 4 1
A
.
B
.
`, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, `1 3 5 2
ABA
BBB
BBA
BAB
ABB
`, 11)
}
