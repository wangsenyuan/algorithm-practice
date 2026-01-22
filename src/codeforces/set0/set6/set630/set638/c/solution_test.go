package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, edges, res := drive(reader)
	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, len(res))
	}
	var cnt int
	marked := make([]bool, n)

	for _, cur := range res {
		cnt += len(cur)
		for _, i := range cur {
			i--
			u, v := edges[i][0], edges[i][1]
			u--
			v--
			if marked[u] || marked[v] {
				t.Fatalf("brigade of (%d %d), is already in use", u+1, v+1)
			}
			marked[u] = true
			marked[v] = true
		}

		for _, i := range cur {
			i--
			u, v := edges[i][0], edges[i][1]
			u--
			v--
			marked[u] = false
			marked[v] = false
		}
	}

	if cnt != n-1 {
		t.Fatalf("Sample result %v, not valid", res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
1 2
3 4
3 2
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
3 4
5 4
3 2
1 3
4 6
`
	expect := 3
	runSample(t, s, expect)
}
