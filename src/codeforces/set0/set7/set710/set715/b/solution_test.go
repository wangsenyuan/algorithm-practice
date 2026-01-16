package main

import (
	"bufio"
	"container/heap"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	n, L, x, y, res := drive(bufio.NewReader(strings.NewReader(s)))

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	adj := make([][]int, n)
	for i, e := range res {
		u, v, w := e[0], e[1], e[2]
		if w == 0 {
			t.Fatalf("Sample result %v, not correct, edge %v is not set", res, e)
		}
		adj[u] = append(adj[u], i)
		adj[v] = append(adj[v], i)
	}
	items := make([]*Item, n)
	pq := make(PriorityQueue, n)
	for i := range n {
		items[i] = new(Item)
		items[i].id = i
		items[i].priority = inf
		items[i].index = i
		pq[i] = items[i]
	}
	items[x].priority = 0

	heap.Init(&pq)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		u := it.id
		for _, i := range adj[u] {
			v := res[i][0] ^ res[i][1] ^ u
			w := res[i][2]
			if items[v].priority > it.priority+w {
				pq.update(items[v], it.priority+w)
			}
		}
	}

	if items[y].priority != L {
		t.Fatalf("Sample result %v, not correct, distance to t is %d, want %d", res, items[y].priority, L)
	}
}

func TestSample1(t *testing.T) {
	s := `5 5 13 0 4
0 1 5
2 1 2
3 2 3
1 4 0
4 3 4
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 1 123456789 0 1
0 1 0
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 1 999999999 1 0
0 1 1000000000
`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 4 13 1 3
1 3 13
2 3 0
2 0 0
1 0 12
`
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5 6 1000000000 0 4
0 1 1
2 0 2
3 0 3
4 1 0
4 2 0
3 4 0
`
	expect := true
	runSample(t, s, expect)
}
