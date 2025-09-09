package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, l, r, res := drive(reader)
	head := -1
	tail := -1
	for i := range n {
		if l[i] != 0 && res[i][0] != l[i] {
			t.Fatalf("Sample result %v, not correct, it can't change the previous element", res)
		}
		if r[i] != 0 && res[i][1] != r[i] {
			t.Fatalf("Sample result %v, not correct, it can't change the next element", res)
		}
		if res[i][0] == 0 {
			if head != -1 {
				t.Fatalf("Sample result %v, not correct, it can't heave more than one heads", res)
			}
			head = i
		}
		if res[i][1] == 0 {
			if tail != -1 {
				t.Fatalf("Sample result %v, not correct, it can't heave more than one tails", res)
			}
			tail = i
		}
	}

	if head < 0 || tail < 0 {
		t.Fatalf("Sample result %v, not correct, it can't have more than one heads or tails", res)
	}

	cnt := 1
	cur := head
	for cur != tail && cnt < n {
		cur = res[cur][1] - 1
		cnt++
	}

	if cur != tail || cnt != n {
		t.Fatalf("Sample result %v, not correct, it can't have more than one heads or tails", res)
	}
}

func TestSample(t *testing.T) {
	runSample(t, `7
4 7
5 0
0 0
6 1
0 2
0 4
1 0
	`)
}
