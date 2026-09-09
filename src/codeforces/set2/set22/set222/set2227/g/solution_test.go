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
	runSample(t, `3
10 20 10
`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
1 1 1 1 1
`, 9)
}

func TestSample3(t *testing.T) {
	runSample(t, `4
5 1 5 1
`, 5)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
100
`, 1)
}
