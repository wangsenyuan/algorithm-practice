package main

import (
	"bufio"
	"cmp"
	"reflect"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, _, edges := process(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	// 看看能否用res恢复edges
	var tmp [][]int
	for _, cur := range res {
		u, v, w := cur[0], cur[1], cur[2]
		tmp = append(tmp, []int{u, v})
		tmp = append(tmp, []int{v, w})
	}
	for _, e := range edges {
		e[0], e[1] = min(e[0], e[1]), max(e[0], e[1])
	}
	for _, e := range tmp {
		e[0], e[1] = min(e[0], e[1]), max(e[0], e[1])
	}
	slices.SortFunc(edges, func(a, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	slices.SortFunc(tmp, func(a, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	if !reflect.DeepEqual(edges, tmp) {
		t.Fatalf("Sample expect %v, but got %v", edges, tmp)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `8 12
1 2
2 3
3 4
4 1
1 3
2 4
3 5
3 6
5 6
6 7
6 8
7 8
	`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 3
1 2
2 3
3 1
	`, false)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 2
1 2
2 3
	`, true)
}
