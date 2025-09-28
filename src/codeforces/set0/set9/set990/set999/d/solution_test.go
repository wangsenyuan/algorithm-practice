package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	m, a, best, res := drive(reader)

	if best != expect {
		t.Errorf("Sample %s, expect best %d, but got %d", s, expect, best)
	}

	n := len(a)
	cnt := make([]int, m)
	var sum int
	for i := range n {
		sum += res[i] - a[i]
		cnt[res[i]%m]++
	}

	if sum != expect {
		t.Fatalf("Sample result %v, not getting the expected best %d, but got %d", res, expect, sum)
	}

	for i := range m {
		if cnt[i]*m != n {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `6 3
3 2 0 6 10 12
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2
0 1 2 3
`
	expect := 0
	runSample(t, s, expect)
}
