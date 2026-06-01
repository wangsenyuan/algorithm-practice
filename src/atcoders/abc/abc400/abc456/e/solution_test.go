package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 4
1 2
1 4
2 4
2 3
3
xxo
xox
oxo
oxx`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 0
4
oooo`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 5
1 4
2 3
4 5
3 4
2 5
7
oxxxxxx
xxoxxxo
xxxoxox
xoxxoxx
oxxxoxx`
	expect := false
	runSample(t, s, expect)
}
