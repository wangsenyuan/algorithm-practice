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
	s := `4 2
abba
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8 1
aabaabaa
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 1
bbabbabbba
`
	expect := 6
	runSample(t, s, expect)
}
