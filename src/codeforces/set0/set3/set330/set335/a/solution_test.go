package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, sheets int) {
	t.Helper()
	var target string
	var n int
	fmt.Fscan(strings.NewReader(input), &target, &n)
	res := drive(bufio.NewReader(strings.NewReader(input)))
	if !validSheet(target, n, sheets, res) {
		t.Fatalf("invalid result %q", res)
	}
}

func validSheet(target string, n int, sheets int, res string) bool {
	res = strings.TrimSpace(res)
	if sheets < 0 {
		return res == "-1"
	}
	lines := strings.Split(res, "\n")
	if len(lines) != 2 {
		return false
	}
	var got int
	if _, err := fmt.Sscan(lines[0], &got); err != nil || got != sheets {
		return false
	}
	sheet := lines[1]
	if len(sheet) != n {
		return false
	}
	need := map[byte]int{}
	for i := range target {
		need[target[i]]++
	}
	have := map[byte]int{}
	for i := range sheet {
		if sheet[i] < 'a' || sheet[i] > 'z' {
			return false
		}
		have[sheet[i]]++
	}
	for c, k := range need {
		if have[c]*got < k {
			return false
		}
	}
	return true
}

func TestSample1(t *testing.T) {
	runSample(t, `banana
4
`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `banana
3
`, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, `banana
2
`, -1)
}
