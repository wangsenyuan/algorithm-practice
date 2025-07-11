package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, n := process(reader)

	k := len(res)

	for _, cur := range res {
		if len(cur) > (2*n+k-1)/k {
			t.Fatalf("cur = %v, n = %d, k = %d", cur, n, k)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 2 1
2 1
3 1`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 4 2
1 2
1 3
1 4
1 5`)
}
