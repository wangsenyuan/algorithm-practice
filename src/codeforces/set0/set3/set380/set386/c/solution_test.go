package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	diversity, ans := process(reader)
	expect_diversity := readNum(reader)
	if diversity != expect_diversity {
		t.Fatalf("expect diversity %d, but got %d", expect_diversity, diversity)
	}
	for _, x := range ans {
		expect := readNum(reader)
		if expect != x {
			t.Fatalf("Sample expect %s, but got %v", s, ans)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `abca
3
4
3
3`)
}

func TestSample2(t *testing.T) {
	runSample(t, `aabacaabbad
4
14
19
28
5`)
}
