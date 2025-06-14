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
	s := `2 2
cb
c b
b a
ab`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `10 10
bbbbbbbbbb
b a
b c
c b
b a
c a
b c
b c
b a
a b
c a
aaaaabbbbb`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `30 20
abcaababcbbcabcbbcabcbabbbbabc
b c
b c
c a
b c
b c
b a
b c
b c
b a
b a
b a
b a
c a
b c
c a
b c
c a
c a
b c
c b
aaaaaaaaaaaaaaabbbabcbabbbbabc`
	runSample(t, s)
}
