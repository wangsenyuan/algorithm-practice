package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res, k, queries := drive(reader)
	n := len(res)
	freq := make([]int, n+1)
	for _, cur := range queries {
		l, r := cur[1]-1, cur[2]-1
		if cur[0] == 1 {
			w := slices.Min(res[l : r+1])
			if w != k {
				t.Fatalf("Sample result %v, is invalid", res)
			}
		} else {
			clear(freq)
			var mex int
			for i := l; i <= r; i++ {
				if res[i] < n {
					freq[res[i]]++
				}
				for freq[mex] > 0 {
					mex++
				}
			}
			if mex != k {
				t.Fatalf("Sample result %v, is invalid", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6 2 2
1 1 3
2 2 6
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 3 1
2 1 3
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 3 2
1 1 1
1 3 3
`)
}

func TestSample4(t *testing.T) {
	runSample(t, `3 2 2
2 1 2
2 2 3
`)
}
