package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	V, a, b, edges, ok, res := process(reader)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}

	if !expect {
		return
	}
	n := len(a)
	g := make([][]bool, n)
	for i := range n {
		g[i] = make([]bool, n)
	}
	for _, e := range edges {
		u, v := e[0], e[1]
		u--
		v--
		g[u][v] = true
		g[v][u] = true
	}
	for _, cur := range res {
		u, v, w := cur[0], cur[1], cur[2]
		u--
		v--
		if !g[u][v] {
			t.Fatalf("Sample result %v, not correct", cur)
		}
		a[u] -= w
		a[v] += w
		if a[u] < 0 || a[v] > V {
			t.Fatalf("Sample result %v, not correct", cur)
		}
	}

	if !reflect.DeepEqual(a, b) {
		t.Fatalf("Sample result not geting correct result %v vs %v", a, b)
	}
}

func TestSample1(t *testing.T) {
	s := `2 10 1
1 9
5 5
1 2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 10 0
5 2
4 2
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 10 0
4 2
4 2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1 1000000000 0
999999999
1000000000
`
	expect := false
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `7 7 5
6 3 0 7 0 7 7
4 5 7 7 7 0 0
1 2
4 3
4 5
4 6
4 7
`
	expect := true
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `6 1 4
0 1 0 1 0 1
1 0 1 0 1 0
1 2
4 3
4 5
4 6
`
	expect := true
	runSample(t, s, expect)
}
