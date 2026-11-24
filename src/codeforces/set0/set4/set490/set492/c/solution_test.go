package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 5 4
5 2
4 7
3 1
3 2
2 5
`, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 5 4
5 2
5 2
`, 0)
}
