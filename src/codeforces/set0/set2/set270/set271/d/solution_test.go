package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `ababab
01000000000000000000000000
1
`, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, `acbacbacaa
00000000000000000000000000
2`, 8)
}
