package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, edges, ok, res := process(reader)
	cnt := readNum(reader)
	if cnt < 0 {
		if ok {
			t.Fatalf("Sample expect %d, but got %t", cnt, ok)
		}
		return
	}
	if !ok {
		t.Fatalf("Sample expect %d, but got %t", cnt, ok)
	}
	// cnt == n - 1
	// expect := readNNums(reader, cnt)
	tmp := []int{0, 0}
	set := NewDSU(n)
	for _, i := range res {
		e := edges[i-1]
		u, v, w := e[0]-1, e[1]-1, e[2]
		tmp[w]++
		if !set.Union(u, v) {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	if set.sz != 1 || tmp[0] != tmp[1] {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1 2
1 1 S
1 1 M
0`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 3
1 2 S
1 3 M
2 3 S
2`)
}

func TestSample3(t *testing.T) {
	runSample(t, `5 6
1 1 S
1 2 M
1 3 S
1 4 M
1 5 M
2 2 S
-1`)
}
