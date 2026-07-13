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
	runSample(t, `7
1101001
3 4 9 100 1 2 3
`, 109)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
10101
3 10 15 15 15
`, 23)
}
