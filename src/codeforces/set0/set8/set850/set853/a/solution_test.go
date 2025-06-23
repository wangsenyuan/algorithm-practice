package main

import (
	"bufio"
	"sort"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	k, c, res, d := process(reader)
	expect := readNum(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}

	var sum int

	n := len(c)
	for i := 0; i < n; i++ {
		if d[i] <= k || d[i] < i+1 {
			t.Fatalf("Sample result %v, not correct", d)
		}
		sum += c[i] * (d[i] - 1 - i)
	}

	sort.Ints(d)

	for i := 1; i < n; i++ {
		if d[i] != d[i-1]+1 {
			t.Fatalf("Sample result %v, not correct", d)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 2
4 2 1 10 2
20`)
}
