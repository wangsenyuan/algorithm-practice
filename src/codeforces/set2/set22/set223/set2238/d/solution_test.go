package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `4
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `8
`, 3)
}

func TestSample4(t *testing.T) {
	runSample(t, `16
`, 4)
}

func TestSample5(t *testing.T) {
	runSample(t, `32
`, 5)
}

func TestSample6(t *testing.T) {
	runSample(t, `67
`, 1)
}

func TestSample7(t *testing.T) {
	runSample(t, `120
`, 7)
}

func TestSample8(t *testing.T) {
	runSample(t, `33
`, 3)
}
