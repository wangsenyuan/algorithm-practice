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
		t.Fatalf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 3 2
3`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 3 3
1`)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 3 2
6`)
}

func TestSample4(t *testing.T) {
	runSample(t, `4 5 2
7`)
}

func TestSample5(t *testing.T) {
	runSample(t, `50 6 3
295630102`)
}
