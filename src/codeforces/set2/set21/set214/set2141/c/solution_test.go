package main

import (
	"strconv"
	"strings"
	"testing"
)

func runSample(t *testing.T, n int) {
	res := solve(n)
	if len(res) > n*(n+2) {
		t.Fatalf("Sample result process %d times, exceeds n * (n+2)", len(res))
	}
	marked := make([][]bool, n+1)
	for i := range n + 1 {
		marked[i] = make([]bool, n+1)
	}
	que := make([]int, 2*n)
	head, tail := n, n
	for _, cur := range res {
		if strings.HasPrefix(cur, "pushback") {
			w := strings.Index(cur, "a[")
			v := strings.Index(cur, "]")
			i, _ := strconv.Atoi(cur[w+2 : v])
			que[head] = i
			head++
		} else if strings.HasSuffix(cur, "pushfront") {
			w := strings.Index(cur, "a[")
			v := strings.Index(cur, "]")
			i, _ := strconv.Atoi(cur[w+2 : v])
			tail--
			que[tail] = i
		} else if cur == "popback" {
			head--
		} else if cur == "popfront" {
			tail++
		} else {
			// min
			if tail == head {
				t.Fatalf("result is invalid, can't query min on empty")
			}
			for i := tail + 1; i < head; i++ {
				if que[i] != que[i-1]+1 {
					t.Fatalf("result is invalid, query min on %v, not continuse", que[tail:head])
				}
			}
			marked[que[tail]][que[head-1]] = true
		}
		if head < tail {
			t.Fatalf("result is invalid")
		}
	}
	for l := 0; l < n; l++ {
		for r := l; r < n; r++ {
			if !marked[l][r] {
				t.Fatalf("Sample result is invalid, it not asked range [%d: %d] (0-based)", l, r)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 3)
}

func TestSample4(t *testing.T) {
	runSample(t, 4)
}

func TestSample5(t *testing.T) {
	runSample(t, 100)
}

