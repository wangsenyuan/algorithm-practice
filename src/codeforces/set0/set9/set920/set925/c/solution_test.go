package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	b, res := drive(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	n := len(b)
	a := make([]int, n)
	a[0] = res[0]
	sum := a[0]
	for i := 1; i < n; i++ {
		sum ^= res[i]
		a[i] = sum
		if a[i-1] >= a[i] {
			t.Fatalf("Sample result %v, not valid", res)
		}
	}

	slices.Sort(b)
	slices.Sort(res)

	if !slices.Equal(b, res) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
1 2 3
`, false)
}

func TestSample2(t *testing.T) {
	runSample(t, `6
4 7 7 12 31 61`, true)
}

func TestSample3(t *testing.T) {
	runSample(t, `10
10 1 1 1 1 1 3 6 7 3`, false)
}
