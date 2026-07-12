package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !strings.EqualFold(res, expect) {
		t.Fatalf("Sample expect %q, but got %q", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `0 0 6 0 6 6 0 6
1 3 3 5 5 3 3 1
`, "Yes")
}

func TestSample2(t *testing.T) {
	runSample(t, `0 0 6 0 6 6 0 6
7 3 9 5 11 3 9 1
`, "No")
}

func TestSample3(t *testing.T) {
	runSample(t, `6 0 6 6 0 6 0 0
7 4 4 7 7 10 10 7
`, "Yes")
}

func TestTouchAtPoint(t *testing.T) {
	runSample(t, `0 0 2 0 2 2 0 2
2 2 3 3 4 2 3 1
`, "Yes")
}
