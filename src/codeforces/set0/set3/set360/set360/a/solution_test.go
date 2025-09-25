package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	ops, res := drive(bufio.NewReader(strings.NewReader(s)))

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	for _, op := range ops {
		o, l, r, d := op[0], op[1], op[2], op[3]
		l--
		r--
		if o == 1 {
			for j := l; j <= r; j++ {
				res[j] += d
			}
		} else {
			var x int = -inf
			for j := l; j <= r; j++ {
				x = max(x, res[j])
			}
			if x != d {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4 5
1 2 3 1
2 1 2 8
2 3 4 7
1 1 3 3
2 3 4 8
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 5
1 2 3 1
2 1 2 8
2 3 4 7
1 1 3 3
2 3 4 13
`
	expect := false
	runSample(t, s, expect)
}
