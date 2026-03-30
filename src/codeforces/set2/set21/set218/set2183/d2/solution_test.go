package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	edges, res := drive(reader)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, len(res))
	}

	n := len(edges) + 1
	assign := make([]int, n)

	for i, cur := range res {
		for _, v := range cur {
			assign[v-1] = i
		}
	}

	adj := make([][]int, n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
		if assign[u] == assign[v] {
			t.Fatalf("Sample result %v, not correct, color(%d) == color(%d)", res, u+1, v+1)
		}
	}

	que := make([]int, n)
	var head, tail int
	que[head] = 0
	head++
	for tail < head {
		var arr []int
		mark := head
		for i := tail; i < mark; i++ {
			u := que[i]
			arr = append(arr, assign[u])
			assign[u] = -1
			for _, v := range adj[u] {
				if assign[v] >= 0 {
					que[head] = v
					head++
				}
			}
		}

		slices.Sort(arr)
		arr = slices.Compact(arr)
		if len(arr) != mark-tail {
			t.Fatalf("Sample result %v is invalid", res)
		}

		tail = mark
	}

}

func TestSample1(t *testing.T) {
	s := `5
3 1
1 2
5 1
4 1`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
3 2
2 4
2 5
1 2`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
3 4
4 1
5 1
1 2`
	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
2 5
3 1
2 1
3 4`
	expect := 3
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5
1 3
1 5
4 3
2 4`
	expect := 3
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `13
2 1
3 2
4 2
5 4
6 3
7 1
8 5
9 6
10 4
11 7
12 8
13 10`
	expect := 3
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `10
5 7
8 1
1 10
2 8
8 4
9 4
6 1
5 3
7 8`
	expect := 4
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `10
7 6
3 7
6 9
7 1
9 8
5 1
3 10
9 2
1 4`
	expect := 4
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := `10
10 6
2 8
4 10
7 5
1 2
7 10
10 9
9 1
7 3`
	expect := 4
	runSample(t, s, expect)
}

func TestSample10(t *testing.T) {
	s := `10
6 8
9 7
4 10
5 9
4 2
3 8
6 5
1 5
1 10`
	expect := 3
	runSample(t, s, expect)
}
