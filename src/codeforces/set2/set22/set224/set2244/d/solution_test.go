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
	runSample(t, `5 3
-1 2 -3 4 -5
1 5 3
`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 2
3 -1 3 -1
4 2
`, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 1
-5 -5 -5
2
`, 5)
}

func TestSample4(t *testing.T) {
	runSample(t, `4 3
3 -1 1 -3
4 3 2
`, 6)
}
