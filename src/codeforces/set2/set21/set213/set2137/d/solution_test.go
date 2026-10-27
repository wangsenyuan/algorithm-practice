package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, possible bool) {
	t.Helper()
	b := parseCase(s)
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if !possible {
		if res != nil {
			t.Fatalf("expect no array, got %v", res)
		}
		return
	}
	if !validArray(b, res) {
		t.Fatalf("invalid array %v for %v", res, b)
	}
}

func parseCase(s string) []int {
	reader := bufio.NewReader(strings.NewReader(s))
	var n int
	fmt.Fscan(reader, &n)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return b
}

func validArray(b []int, a []int) bool {
	n := len(b)
	if len(a) != n {
		return false
	}
	freq := make([]int, n+1)
	for _, v := range a {
		if v < 1 || v > n {
			return false
		}
		freq[v]++
	}
	for i, v := range a {
		if freq[v] != b[i] {
			return false
		}
	}
	return true
}

func TestSample1(t *testing.T) {
	runSample(t, `4
1 2 3 4
`, false)
}

func TestSample2(t *testing.T) {
	runSample(t, `6
1 2 2 3 3 3
`, true)
}

func TestSample3(t *testing.T) {
	runSample(t, `6
6 6 6 6 6 6
`, true)
}
