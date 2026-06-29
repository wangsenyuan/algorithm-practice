package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "11 4 2\nabbaaabaaba\n", 3)
}

func TestSample2(t *testing.T) {
	runSample(t, "13 1 2\nbbbbbbbbbbbbb\n", 0)
}

func TestSample3(t *testing.T) {
	// a, aa, aaa, aaaa, aaaaa
	runSample(t, "5 1 1\naaaaa\n", 15)
}

func TestSample4(t *testing.T) {
	runSample(t, "5 1 1\nabbbb\n", 1)
}

func TestSample5(t *testing.T) {
	runSample(t, "5 1 1\nbbbbb\n", 0)
}
