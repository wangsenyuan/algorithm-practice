package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	expect := readNum(reader)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 3 1
aba
7`)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 1 1
abcd
4
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 10 1
aaaa
12
`)
}
