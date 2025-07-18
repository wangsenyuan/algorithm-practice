package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readString(reader)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
cbabc
a`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `2
abcab
aab`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `3
bcabcbaccba
aaabb`
// 3 6 7 a
	runSample(t, s)
}
