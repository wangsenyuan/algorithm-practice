package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `10
3 10 7 10 7 6 7 6 5 14
`, 7)
}

func TestSample2(t *testing.T) {
	runSample(t, `6
210 210 210 210 210 210
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `21
49 30 50 21 35 15 21 70 35 9 50 70 21 49 30 50 70 15 9 21 30
`, 34)
}
