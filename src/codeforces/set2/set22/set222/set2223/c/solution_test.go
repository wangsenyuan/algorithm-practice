package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	t.Helper()
	t.Skip("solve TODO")
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 1
1 1
10 20
4
`, "2")
}

func TestSample2(t *testing.T) {
	runSample(t, `10 5
1 2 2 2 1 1 3 4 5
1 2 3 4 5 6 7 8 9
1 2 3 4 5
`, "6 7 9 6 7")
}
