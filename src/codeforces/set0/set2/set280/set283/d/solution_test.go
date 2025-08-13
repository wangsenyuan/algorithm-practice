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
	s := `3
6 4 1
`
	runSample(t, s, 0)
}


func TestSample2(t *testing.T) {
	s := `4
20 6 3 4
`
	runSample(t, s, 2)
}
