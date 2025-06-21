package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	for _, ans := range res {
		expect := readString(reader)
		if len(ans) != len(expect) {
			t.Fatalf("Sample expect %s, but got %s", expect, ans)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
123456789
100000000
100123456
9
000
01`)
}
