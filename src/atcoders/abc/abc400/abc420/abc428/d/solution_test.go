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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 80
`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `183 5000
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `18 10
`, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, `824 5000000000`, 1421)
}
