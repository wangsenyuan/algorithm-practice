package main

import (
	"bufio"
	"slices"
	"sort"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	a, res := process(bufio.NewReader(strings.NewReader(s)))

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	i, j := res[0]-1, res[1]-1
	a[i], a[j] = a[j], a[i]
	if sort.IntsAreSorted(a) {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	slices.Reverse(a)

	if sort.IntsAreSorted(a) {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
1 2 3 4
`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 1 1
`, false)
}

func TestSample3(t *testing.T) {
	runSample(t, `3
1 2 2
`, true)
}


func TestSample4(t *testing.T) {
	runSample(t, `3
2 2 1
`, true)
}

func TestSample5(t *testing.T) {
	runSample(t, `3
2 1 2
`, false)
}
