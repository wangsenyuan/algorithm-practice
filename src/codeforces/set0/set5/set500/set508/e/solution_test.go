package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	segments, res := drive(reader)
	if (res != "IMPOSSIBLE") != expect {
		t.Fatalf("Sample expect %t, but got %s", expect, res)
	}
	if !expect {
		return
	}

	n := len(segments)
	stack := make([]int, n)
	var dist []int
	var top int
	pos := make([]int, len(res))
	for i := range res {
		if res[i] == '(' {
			pos[i] = len(dist)
			dist = append(dist, 0)
			stack[top] = i
			top++
		} else {
			if top == 0 {
				t.Fatalf("Sample result %s, not a regular bracket sequence", res)
			}
			j := stack[top-1]
			dist[pos[j]] = i - j
			top--
		}
	}

	if top > 0 {
		t.Fatalf("Sample result %s, not a regular bracket sequence", res)
	}

	for i, cur := range segments {
		l, r := cur[0], cur[1]
		if dist[i] < l || dist[i] > r {
			t.Fatalf("Sample result voliates contraints at %d", i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4
1 1
1 1
1 1
1 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
5 5
3 3
1 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
5 5
3 3
2 2
`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3
2 3
1 4
1 4
`
	expect := true
	runSample(t, s, expect)
}
