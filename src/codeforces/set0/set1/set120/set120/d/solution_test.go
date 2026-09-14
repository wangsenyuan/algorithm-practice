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
	runSample(t, `3 3
1 1 1
1 1 1
1 1 1
3 3 3
`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 5
1 1 1 1 1
2 2 2 2 2
3 6 6
`, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 3
1 2 3
3 1 2
2 3 1
5 6 7
`, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, `3 3
0 0 0
0 0 1
1 1 0
2 1 0
`, 1)
}

func TestSample5(t *testing.T) {
	runSample(t, `3 3
0 0 0
0 1 0
0 0 0
1 0 0
`, 2)
}
