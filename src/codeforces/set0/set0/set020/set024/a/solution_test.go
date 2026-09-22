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
	runSample(t, `3
1 3 1
1 2 1
3 2 1
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 3 1
1 2 5
3 2 1
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `6
1 5 4
5 3 8
2 4 15
1 6 16
2 3 23
4 6 42
`, 39)
}

func TestSample4(t *testing.T) {
	runSample(t, `4
1 2 9
2 3 8
3 4 7
4 1 5
`, 0)
}
