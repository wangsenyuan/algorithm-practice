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
	if res != expect {
		t.Fatalf("Sample expect %q, but got %q", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
5 4 3 3 2 5 3 5 3
`, "95")
}

func TestSample2(t *testing.T) {
	runSample(t, `20
1 1 1 1 1 1 1 1 1
`, "99999999999999999999")
}
