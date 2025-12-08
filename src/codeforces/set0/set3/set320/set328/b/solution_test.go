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
	s := `42
23454
	`
	runSample(t, s, 2)
}

func TestSample2(t *testing.T) {
	s := `169
12118999
	`
	runSample(t, s, 1)
}
