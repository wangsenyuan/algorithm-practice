package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, length int) {
	t.Helper()
	var word string
	fmt.Fscan(strings.NewReader(input), &word)
	res := drive(bufio.NewReader(strings.NewReader(input)))
	if len(res) != length || !subsequence(word, res) || hasTypo(res) {
		t.Fatalf("invalid result %q", res)
	}
}

func subsequence(s string, t string) bool {
	j := 0
	for i := 0; i < len(s) && j < len(t); i++ {
		if s[i] == t[j] {
			j++
		}
	}
	return j == len(t)
}

func hasTypo(t string) bool {
	for i := 0; i+2 < len(t); i++ {
		if t[i] == t[i+1] && t[i] == t[i+2] {
			return true
		}
	}
	for i := 0; i+3 < len(t); i++ {
		if t[i] == t[i+1] && t[i+2] == t[i+3] {
			return true
		}
	}
	return false
}

func TestSample1(t *testing.T) {
	runSample(t, `helloo
`, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, `woooooow
`, 4)
}
