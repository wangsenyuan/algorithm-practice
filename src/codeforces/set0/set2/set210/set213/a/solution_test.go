package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
1
0
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
2 2 1 1 3
1 5
2 5 1
2 5 4
1 5
0`, 7)
}
