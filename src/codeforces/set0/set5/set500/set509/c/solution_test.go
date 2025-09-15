package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))

	b, a := drive(reader)

	n := len(b)
	if a[n-1] != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, a[n-1])
	}

	for i := range n {
		if digit_sum(a[i]) != b[i] {
			t.Fatalf("Sample result %v, not having the correct digit sum", a)
		}
		if i > 0 && cmp(a[i-1], a[i]) >= 0 {
			t.Fatalf("Sample result %v, not correct", a)
		}
	}
}

func digit_sum(s string) int {
	var res int
	for i := range len(s) {
		res += int(s[i] - '0')
	}
	return res
}

func cmp(a, b string) int {
	if len(a) > len(b) {
		return 1
	}
	if len(a) < len(b) {
		return -1
	}
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

func TestSample1(t *testing.T) {
	s := `3
1
2
3
`
	expect := "3"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
3
2
1
`
	expect := "100"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
8
8
5
1
2
7
3
8
9
4
`
	expect := "121"
	runSample(t, s, expect)
}