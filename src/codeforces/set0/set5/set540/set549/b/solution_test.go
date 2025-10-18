package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, contacts, a, res := drive(reader)

	deg := make([]int, n)

	for _, u := range res {
		for j := range n {
			if contacts[u-1][j] == '1' {
				deg[j]++
			}
		}
	}

	for i := range n {
		if deg[i] == a[i] {
			t.Fatalf("Sample result %v, leading to deg[%d] = a[%d]", res, i, i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3
101
010
001
0 1 2
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `1
1
1
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `4
1111
0101
1110
0001
1 0 1 0
`
	runSample(t, s)
}
