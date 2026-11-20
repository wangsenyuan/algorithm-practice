package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 7
3 5 6
0 3 4
3 3 3
`, 31)
}

func TestSample2(t *testing.T) {
	runSample(t, `1 45000
40000 50000 100000
`, 4500000000)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 100
2 10 50
50 60 16
20 21 25
`, -1)
}

func TestSample4(t *testing.T) {
	runSample(t, `4 12
2 5 1
1 2 2
2 3 7
3 10 4
`, 35)
}
