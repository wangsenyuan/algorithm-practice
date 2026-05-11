package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 3
0 1 2 0
`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `8 4
0 1 2 3 1 2 0 1
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `9 5
1 0 1 3 4 3 2 1 0
`, 1920)
}

func TestSample4(t *testing.T) {
	runSample(t, `15 14
3 0 1 2 3 4 5 6 7 8 9 10 11 12 13
`, 138007136)
}

func TestSample5(t *testing.T) {
	runSample(t, `5 5
4 3 0 1 2
`, 8)
}

func TestSample6(t *testing.T) {
	runSample(t, `5 2
0 1 1 1 0
`, 0)
}

func TestSample7(t *testing.T) {
	runSample(t, `3 2
0 1 1
`, 0)
}
