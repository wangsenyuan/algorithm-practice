package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans, m, w := process(reader)
	if (len(ans) == m) != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, ans)
	}
	if !expect {
		return
	}
	var l, r int
	for i := 0; i < m; i++ {
		v := ans[i]
		if w[v-1] == '0' {
			t.Fatalf("Sample result %v, not valid", ans)
		}
		if i&1 == 0 {
			l += v
			if l <= r {
				t.Fatalf("Sample result %v, not valid", ans)
			}
		} else {
			r += v
			if l >= r {
				t.Fatalf("Sample result %v, not valid", ans)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `0000000101
3`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `1000000000
2`, false)
}
