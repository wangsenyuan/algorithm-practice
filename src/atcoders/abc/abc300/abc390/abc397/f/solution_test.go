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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
3 1 4 1 5
`, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, `10
2 5 6 4 4 1 1 3 1 4
`, 9)
}
