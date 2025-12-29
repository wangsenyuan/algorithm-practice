package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	coupons, best, ans := drive(reader)

	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}

	l, r := -1<<60, (1 << 60)

	n := len(coupons)
	marked := make([]bool, n)
	for _, i := range ans {
		marked[i-1] = true
	}

	for i, cur := range coupons {
		if marked[i] {
			l = max(l, cur[0])
			r = min(r, cur[1])
		}
	}

	if max(0, r-l+1) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, max(0, r-l+1))
	}
}

func TestSample1(t *testing.T) {
	s := `4 2
1 100
40 70
120 130
125 180
`
	expect := 31
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 2
1 12
15 20
25 30
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 2
1 10
5 15
14 50
30 70
99 100
`
	expect := 21
	runSample(t, s, expect)
}
