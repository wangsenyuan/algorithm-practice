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
	runSample(t, "5071\n", 4)
}

func TestSample2(t *testing.T) {
	runSample(t, "705\n", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, "1241367\n", -1)
}

func TestSample4(t *testing.T) {
	runSample(t, "17010\n", 1)
}

func TestSample5(t *testing.T) {
	runSample(t, "50267\n", 5)
}
