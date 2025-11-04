package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, m, l, res := drive(reader)

	if len(res) != expect {
		t.Fatalf("Sample expect %d nums, but got %v", expect, res)
	}

	if l > m {
		t.Fatalf("Sample expect l <= m, but got %d > %d", l, m)
	}

	w := 1
	for _, i := range res {
		w = lcm(w, a[i-1])
	}

	if w != l {
		t.Fatalf("Sample expect w == l, but got %d != %d", w, l)
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	c := gcd(a, b)
	return a / c * b
}

func TestSample1(t *testing.T) {
	s := `7 8
6 2 9 2 7 2 3
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 4
2 2 2 3 3 3
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 1
2
`
	expect := 0
	runSample(t, s, expect)
}
