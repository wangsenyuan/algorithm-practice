package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	p, q, ans := process(reader)

	if len(ans) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, ans)
	}
	if !expect {
		return
	}
	for _, i := range q {
		if ans[i-1] != -p[i-1] {
			t.Fatalf("Sample result not correct, it must be negative at position %d", i)
		}
	}

	n := len(p)
	stack := make([]int, n)
	var top int
	for i, v := range p {
		if ans[i] > 0 {
			stack[top] = v
			top++
		} else {
			// ans[i] = -1
			if top == 0 || stack[top-1] != v {
				t.Fatalf("Sample result %v, not correct", ans)
			}
			top--
		}
	}

	if top != 0 {
		t.Fatalf("Sample result %v, not correct", ans)
	}
}

func TestSample1(t *testing.T) {
	s := `2
1 1
0
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
1 1 1 1
1 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
1 2 2 1
2 3 4
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `8
1 1 2 1 1 2 1 1
2 7 8
`
	expect := true
	runSample(t, s, expect)
}
