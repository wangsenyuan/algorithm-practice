package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func mex(vals []int) int {
	seen := make([]bool, len(vals)+1)
	for _, v := range vals {
		if 0 <= v && v < len(seen) {
			seen[v] = true
		}
	}
	for i := range seen {
		if !seen[i] {
			return i
		}
	}
	return len(seen)
}

func groups(a []int, labels string) (A, B, C []int) {
	for i, ch := range labels {
		switch ch {
		case 'A':
			A = append(A, a[i])
		case 'B':
			B = append(B, a[i])
		case 'C':
			C = append(C, a[i])
		}
	}
	return
}

func validAssignment(a []int, labels string) bool {
	if len(labels) != len(a) {
		return false
	}
	for _, ch := range labels {
		if ch != 'A' && ch != 'B' && ch != 'C' {
			return false
		}
	}
	A, B, C := groups(a, labels)
	return mex(A) == mex(B) && mex(C) == 0
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

func runSample(t *testing.T, s string, expectOK bool) {
	t.Helper()
	t.Skip("solve TODO")
	a := parseCase(s)
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !expectOK {
		if res != "NO" {
			t.Fatalf("Sample expect NO, but got %v", res)
		}
		return
	}
	lines := strings.Split(strings.TrimSpace(res), "\n")
	if len(lines) != 2 || strings.ToUpper(lines[0]) != "YES" {
		t.Fatalf("Sample expect YES and an assignment, but got %v", res)
	}
	if !validAssignment(a, strings.TrimSpace(lines[1])) {
		t.Fatalf("invalid assignment %q for %v", lines[1], a)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6
1 0 0 1 2 1
`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `4
0 0 0 0
`, true)
}

func TestSample3(t *testing.T) {
	runSample(t, `3
0 2 2
`, false)
}

func TestSample4(t *testing.T) {
	runSample(t, `4
6 7 6 7
`, true)
}

func TestSample5(t *testing.T) {
	runSample(t, `5
0 0 0 1 2
`, true)
}
