package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `100 10
a
aaaaaaaaaa
`, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, `1 1
abacaba
abzczzz
`, 4)
}
