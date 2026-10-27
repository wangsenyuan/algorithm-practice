package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1029 102
`, 999)
}

func TestSample2(t *testing.T) {
	runSample(t, `27191 17
`, 27189)
}

func TestSample3(t *testing.T) {
	runSample(t, `1000000000000000000 999999999999999999
`, 999999999999999999)
}
