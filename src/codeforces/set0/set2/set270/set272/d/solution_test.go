package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
1
2
7
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
1 2
2 3
11
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
5
5
100
`, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, `3
1 1 1
1 1 1
1000000007
`, 90)
}

func TestSample5(t *testing.T) {
	runSample(t, `2
1 2
1 1
1000
`, 3)
}
