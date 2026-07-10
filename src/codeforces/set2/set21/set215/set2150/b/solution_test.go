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
	if len(res) != 1 || res[0] != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
5
2 2 1 0 0
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `1
2
2 0
`, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
2
1 1
`, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
4
3 1 0 0
`, 2)
}

func TestSample5(t *testing.T) {
	runSample(t, `1
4
0 0 0 0
`, 0)
}
