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
	s := `2 2
1 2 4
2 1 4
1`
	runSample(t, s, 16)
}

func TestSample2(t *testing.T) {
	s := `3 3
1 2 4
2 3 3
1 3 8
1`
	runSample(t, s, 8)
}
