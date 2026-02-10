package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %q, but got %q", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "2 2 1 1", "4774")
}

func TestSample2(t *testing.T) {
	runSample(t, "4 7 3 1", "-1")
}

func TestSample3(t *testing.T) {
	runSample(t, "1 1 1 1", "-1")
}

func TestSample4(t *testing.T) {
	runSample(t, "2 2 1 2", "7474")
}

func TestSample5(t *testing.T) {
	runSample(t, "1 7 1 1", "74777777")
}