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
	s := `3 5
..+.#
#+..+
+.#+.
14
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4 10
...+.##+.+
+#++..+++#
++.#++++..
.+##.++#.+
42
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `2 2
..
++
-1
`
	runSample(t, s)
}
