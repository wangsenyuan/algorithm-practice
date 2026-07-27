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
		t.Errorf("Sample expect %q, but got %q", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2
`, "white\n1 2")
}

func TestSample2(t *testing.T) {
	runSample(t, `3
`, "black")
}
