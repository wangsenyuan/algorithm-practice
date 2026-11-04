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
	runSample(t, `5 5
1 2
2 2 3
2 3 4
2 4 5
1 5
`, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, `8 7
0
3 1 2 3
1 1
2 5 4
2 6 7
1 3
2 7 4
1 1
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 2
1 2
0
`, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, `2 2
	0
	0
`, 2)
}
