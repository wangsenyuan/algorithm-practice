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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
0 1 10
0 2 20
`, 15)
}

func TestSample2(t *testing.T) {
	runSample(t, `4
0 1 3
0 2 9
0 3 27
`, 13)
}

func TestSample3(t *testing.T) {
	runSample(t, `7
0 1 3
0 5 7
1 2 2
1 3 1
1 4 5
5 6 8
`, 400000019)
}

func TestSample4(t *testing.T) {
	runSample(t, `11
1 0 6646
2 0 8816
3 2 9375
4 2 5950
5 1 8702
6 2 2657
7 2 885
8 7 2660
9 2 5369
10 6 3798
`, 153869806)
}

func TestSample5(t *testing.T) {
	runSample(t, `6
0 1 8
0 2 24
1 3 40
1 4 16
4 5 8
`, 39)
}
