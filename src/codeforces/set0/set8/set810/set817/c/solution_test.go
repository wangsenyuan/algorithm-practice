package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "12 1", 3)
}

func TestSample2(t *testing.T) {
	runSample(t, "25 20", 0)
}

func TestSample3(t *testing.T) {
	runSample(t, "10 9", 1)
}
