package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 200000
3 -1 9 -5 4
0 0 0 0 0
`, 11)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 5
10 10 10 10
1 1 1 1
`, 41)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 1
2 -7 3
1 11 3
`, 9)
}

func TestSample4(t *testing.T) {
	runSample(t, `3 2
2 -7 3
1 11 3
`, 3)
}

func TestSample5(t *testing.T) {
	runSample(t, `1 1
-3
2
`, -1)
}
