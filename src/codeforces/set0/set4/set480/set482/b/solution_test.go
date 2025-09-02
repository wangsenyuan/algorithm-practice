package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, queries, ans := drive(reader)
	if len(ans) == n != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, ans)
	}

	if !expect {
		return
	}

	for _, cur := range queries {
		l, r, d := cur[0]-1, cur[1]-1, cur[2]
		tmp := 1<<H - 1
		for i := l; i <= r; i++ {
			tmp &= ans[i]
		}

		if tmp != d {
			t.Fatalf("Sample expect %d, but got %d", d, tmp)
		}
	}

}

func TestSample1(t *testing.T) {
	s := `3 1
1 3 3`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 2
1 3 3
1 3 2
`
	expect := false
	runSample(t, s, expect)
}
