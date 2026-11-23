package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect string) {
	t.Helper()

	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2
ab
cd
`, "DRAW")
}

func TestSample2(t *testing.T) {
	runSample(t, `2
xa
ay
`, "FIRST")
}

func TestSample3(t *testing.T) {
	runSample(t, `3
aab
bcb
bac
`, "DRAW")
}
