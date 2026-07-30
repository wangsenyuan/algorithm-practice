package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 2
1 2
2 3
`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 3
1 2 3
2 3 1
`, false)
}

func TestSample3(t *testing.T) {
	runSample(t, `10 1
1
1
`, false)
}

func TestSample4(t *testing.T) {
	runSample(t, `7 8
1 6 2 7 5 4 2 2
3 2 7 2 1 2 3 3
`, true)
}
