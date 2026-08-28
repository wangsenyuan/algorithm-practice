package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	t.Helper()
	t.Skip("solve TODO")
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 1
1 7
3 5
`, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 2
1 5 3
6 2 4
`, 9)
}

func TestSample3(t *testing.T) {
	runSample(t, `5 4
1 16 10 10 16
3 2 2 15 15
`, 30)
}

func TestSample4(t *testing.T) {
	runSample(t, `4 1
23 1 18 4
19 2 10 3
`, 16)
}

func TestSample5(t *testing.T) {
	runSample(t, `10 10
4 3 2 100 4 1 2 4 5 5
1 200 4 5 6 1 10 2 3 4
`, 312)
}
