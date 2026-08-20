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
	runSample(t, `2 3
10 20
2 2 2
5 5 5
`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 1
100 100 100
1
2
3
`, 1)
}
