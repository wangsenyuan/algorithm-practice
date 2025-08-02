package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	res, a, n, m := process(reader)

	if len(res) > 0 != expect || expect && len(res) != m {
		t.Fatalf("Sample expect %t, but got %d", expect, len(res))
	}
	if !expect {
		return
	}

	x := getDist(n, res, a)

	b := make([]int, n)

	for i := range n {
		b[i] = i + 1
	}

	y := getDist(n, res, b)

	if reflect.DeepEqual(x, y) {
		t.Fatal("Sample resul not correct")
	}

	for i := range n {
		for j := range n {
			if y[i][j] == inf {
				t.Fatalf("Sample result graph is not connected")
			}
		}
	}
}

const inf = 1 << 60

func getDist(n int, edges [][]int, a []int) [][]int {
	dist := make([][]int, n)
	for i := range n {
		dist[i] = make([]int, n)
		for j := range n {
			dist[i][j] = inf
		}
		dist[i][i] = 0
	}
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		dist[u][v] = 1
		dist[v][u] = 1
	}

	for _, v := range a {
		v--
		for i := range n {
			for j := range n {
				dist[i][j] = min(dist[i][j], dist[i][v]+dist[v][j])
			}
		}
	}
	return dist
}

func TestSample1(t *testing.T) {
	s := `3 2 2
1 2
`
	expect := true

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3 2
1 2
`
	expect := false

	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `300 43056 2
5 6
`
	expect := true

	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `300 44849 2
1 300
`
	expect := true

	runSample(t, s, expect)
}
