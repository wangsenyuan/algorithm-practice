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
	runSample(t, `2
1 0
2 1 2
`, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 1
2 2 3
3 4 5 6
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
4 1 7 8 10
2 5 6
2 0 7
2 6 6
2 6 8
`, 50)
}

func TestSample4(t *testing.T) {
	runSample(t, `2
1 3
3 0 1 2
`, 8)
}

func TestSample5(t *testing.T) {
	runSample(t, `2
6 0 0 1 2 2 3
3 0 2 3
`, 43)
}

func TestSample6(t *testing.T) {
	runSample(t, `10
1 0
9 7 8 0 1 5 6 4 3 2
8 4 3 8 6 2 5 0 1
7 2 3 0 1 0 4 0
2 3 1
9 2 0 5 4 1 3 0 0 0
7 6 3 2 4 1 8 0
5 3 2 4 1 0
4 0 3 1 1
3 0 3 2
`, 19202)
}