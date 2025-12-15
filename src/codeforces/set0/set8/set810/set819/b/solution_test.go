package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	p, res := drive(reader)

	if res[0] != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	k := res[1]

	n := len(p)
	if k != 0 {
		slices.Reverse(p[:n-k])
		slices.Reverse(p[n-k:])
		slices.Reverse(p)
	}

	var sum int
	for i := range n {
		sum += abs(i + 1 - p[i])
	}

	if sum != expect {
		t.Fatalf("Sample result %v, not getting the expected sum %d, but got %d", res, expect, sum)
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 2 3
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
2 3 1
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
3 2 1
`
	// k = 1, 1 3 2, 变化 = -1 + 1 - 2 = -2
	expect := 2
	runSample(t, s, expect)
}


func TestSample4(t *testing.T) {
	s := `10
10 1 9 2 8 3 7 4 6 5
`
	// k = 1, 1 3 2, 变化 = -1 + 1 - 2 = -2
	expect := 24
	runSample(t, s, expect)
}
