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
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "2 1 1", "BitLGM")
}

func TestSample2(t *testing.T) {
	runSample(t, "2 1 2", "BitAryo")
}

func TestSample3(t *testing.T) {
	runSample(t, "3 1 2 1", "BitLGM")
}

func TestSample4(t *testing.T) {
	runSample(t, "2 3 5", "BitAryo")
}

func TestSample5(t *testing.T) {
	// 102 184 222
	runSample(t, "3 184 222 102", "BitAryo")
}
