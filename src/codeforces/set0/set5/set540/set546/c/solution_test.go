package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Fatalf("Sample expect %q, but got %q", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
2 1 3
2 4 2
`, "6 2")
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 2
2 1 3
`, "-1")
}

func TestSample3(t *testing.T) {
	runSample(t, `4
3 2 3 1
1 4
`, "7 2")
}

func TestSample4(t *testing.T) {
	runSample(t, `5
4 4 1 3 2
1 5
`, "6 2")
}
