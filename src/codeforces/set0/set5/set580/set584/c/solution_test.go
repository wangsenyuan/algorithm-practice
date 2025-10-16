package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	s1, s2, x, res := drive(reader)
	if res != "-1" != expect {
		t.Fatalf("Sample expect %t, but got %s", expect, res)
	}
	if !expect {
		return
	}

	check := func(a string, b string) int {
		var cnt int
		for i := range len(a) {
			if a[i] != b[i] {
				cnt++
			}
		}
		return cnt
	}

	if check(s1, res) != x || check(s2, res) != x {
		t.Fatalf("Sample result %s", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 2
abc
xyc`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `1 0
c
b
`, false)
}
