package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	t.Helper()
	a := readNums(s)
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if !valid(a, res) {
		t.Fatalf("Sample %v got invalid coloring %v", a, res)
	}
}

func readNums(s string) []int {
	reader := bufio.NewReader(strings.NewReader(s))
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return a
}

func valid(a []int, colors []int) bool {
	if len(colors) != len(a) || len(a) == 0 {
		return false
	}
	m := 0
	used := map[int]bool{}
	for _, c := range colors {
		if c < 1 || c > 11 {
			return false
		}
		used[c] = true
		if c > m {
			m = c
		}
	}
	if len(used) != m {
		return false
	}
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			if colors[i] == colors[j] && gcd(a[i], a[j]) == 1 {
				return false
			}
		}
	}
	return true
}

func TestSample1(t *testing.T) {
	runSample(t, `3
6 10 15
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
4 9
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `23
437 519 865 808 909 391 194 291 237 395 323 365 511 497 781 737 871 559 731 697 779 841 961
`)
}
