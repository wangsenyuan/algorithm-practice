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
`, "1/1")
}

func TestSample2(t *testing.T) {
	runSample(t, `3
`, "8/3")
}

func TestSample3(t *testing.T) {
	runSample(t, `4
`, "2/1")
}

func TestSample4(t *testing.T) {
	runSample(t, `5
`, "18/5")
}
