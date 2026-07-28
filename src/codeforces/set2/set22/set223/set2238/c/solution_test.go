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
	runSample(t, `5
1 2 3 4
`, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 1
`, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, `7
1 2 1 3 5 5
`, 9)
}

func TestSample4(t *testing.T) {
	runSample(t, `10
1 1 3 2 2 4 4 4 3
`, 15)
}

func TestSample5(t *testing.T) {
	runSample(t, `15
1 2 1 3 3 4 3 7 3 10 6 7 1 9
`, 22)
}
