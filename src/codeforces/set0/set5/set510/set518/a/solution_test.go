package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	x, y, res := process(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %s", expect, res)
	}

	if !expect {
		return
	}
	if res <= x || y <= res {
		t.Fatalf("Sample result %s, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `abcdefg
abcdefh
`, false)
}

func TestSample2(t *testing.T) {
	runSample(t, `vklldrxnfgyorgfpfezvhbouyzzzzz
vklldrxnfgyorgfpfezvhbouzaaadv
`, true)
}
