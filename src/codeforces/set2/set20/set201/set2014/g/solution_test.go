package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1 1 3
1 5`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 3 3
1 5
2 7`, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 5 2
1 9
2 6
4 9
5 6`, 4)
}

func TestSample4(t *testing.T) {
	runSample(t, `5 2 4
4 7
5 3
7 1
11 2
12 1`, 5)
}

func TestSample5(t *testing.T) {
	runSample(t, `4 1 3
5 10
9 4
14 8
15 3`, 10)
}

func TestSample6(t *testing.T) {
	runSample(t, `5 5 5
8 9
10 7
16 10
21 5
28 9
`, 6)
}
