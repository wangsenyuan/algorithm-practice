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
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `10 3 4 4 3
`, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, `10 3 4 4 5
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `100000 100000 100000 100000 1
`, 100000)
}

func TestSample4(t *testing.T) {
	runSample(t, `5 4 5 3 3
`, 5)
}

func TestSample5(t *testing.T) {
	runSample(t, `100 9 6 3 6
`, 7)
}
