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
	runSample(t, `6 7 5 4
5 5 2 2 5
1 3 3
1 5 2
1 6 5
2 5 4
2 6 7
3 4 11
3 5 3`, 3)
}
