package main

import (
	"bufio"
	"strings"
	"testing"
)

const inf = 1 << 60

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	workers, res := drive(reader)
	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, len(res))
	}
	l, r := inf, -inf
	for _, i := range res {
		i--
		l = min(l, workers[i][1])
		r = max(r, workers[i][1])
	}
	// l是最小的技能，r是最大的技能
	for _, i := range res {
		cur := workers[i-1]
		if cur[0] > l || cur[2] < r {
			t.Fatalf("Sample result %v is invalid", res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
2 8 9
1 4 7
3 6 8
5 8 10
`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `6
3 5 16
1 6 11
4 8 12
7 9 16
2 10 14
8 13 15
`, 4)
}
