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
		t.Errorf("expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 4 3
1 1 1
1 2 2
2 1 3
2 2 7
`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 4 7
1 1 1
1 2 2
2 1 3
2 2 7
`, 8)
}
