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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
1 2 3 4 5
ABABA
`, 11)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
1 2 3 4 5
AAAAA
`, 15)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
1
B
`, 1)
}
