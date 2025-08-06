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
	runSample(t, `5
1 2 2 1 2
`, 120)
}

func TestSample2(t *testing.T) {
	runSample(t, `8
1 2 2 1 2 1 1 2
`, 16800)
}
