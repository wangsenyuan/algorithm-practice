package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	t.Helper()
	a := parseCase(s)
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if !validStrings(a, res) {
		t.Fatalf("invalid strings %q for %v", res, a)
	}
}

func parseCase(s string) []int {
	reader := bufio.NewReader(strings.NewReader(s))
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return a
}

func validStrings(a []int, ss []string) bool {
	if len(ss) != len(a)+1 {
		return false
	}
	for _, s := range ss {
		if len(s) < 1 || len(s) > 200 {
			return false
		}
		for _, ch := range s {
			if ch < 'a' || ch > 'z' {
				return false
			}
		}
	}
	for i, need := range a {
		if lcp(ss[i], ss[i+1]) != need {
			return false
		}
	}
	return true
}

func lcp(s, t string) int {
	n := min(len(s), len(t))
	i := 0
	for i < n && s[i] == t[i] {
		i++
	}
	return i
}

func TestSample1(t *testing.T) {
	runSample(t, `4
1 2 4 2
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
5 3
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `3
1 3 1
`)
}

func TestSample4(t *testing.T) {
	runSample(t, `3
0 0 0
`)
}

func TestSample5(t *testing.T) {
	runSample(t, `3
1 1 0
`)
}

func TestSample6(t *testing.T) {
	runSample(t, `3
2 1 0
`)
}
