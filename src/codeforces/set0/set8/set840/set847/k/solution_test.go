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
	s := `3 5 3 1 8
BerBank University
University BerMall
University BerBank
	`
	runSample(t, s, 11)
}

func TestSample2(t *testing.T) {
	s := `4 2 1 300 1000
a A
A aa
aa AA
AA a
	`
	runSample(t, s, 5)
}

func TestSample3(t *testing.T) {
	s := `2 3 1 4 6
AaCdC CdD
aBACc CdD
	`
	runSample(t, s, 6)
}
