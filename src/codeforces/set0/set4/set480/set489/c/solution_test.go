package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input, expect string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "2 15\n", "69 96")
}

func TestSample2(t *testing.T) {
	runSample(t, "3 0\n", "-1 -1")
}

func TestSample3(t *testing.T) {
	runSample(t, "100 100\n", "1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000099999999999 9999999999910000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
}

func TestSample4(t *testing.T) {
	runSample(t, "1 0\n", "0 0")
}
