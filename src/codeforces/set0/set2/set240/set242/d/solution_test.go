package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	// 假设始终有答案
	reader := bufio.NewReader(strings.NewReader(s))
	n, edges, a, res := process(reader)

	marked := make([]bool, n)
	cnt := make([]int, n)

	for _, x := range res {
		marked[x-1] = true
		cnt[x-1]++
	}

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		if marked[u] {
			cnt[v]++
		}
		if marked[v] {
			cnt[u]++
		}
	}

	for i := range n {
		if a[i] == cnt[i] {
			t.Fatalf("Sample result %v, it has same value at %d", res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5 5
2 3
4 1
1 5
5 3
2 1
1 1 2 0 2
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4 2
1 2
3 4
0 0 0 0
`
	runSample(t, s)
}

