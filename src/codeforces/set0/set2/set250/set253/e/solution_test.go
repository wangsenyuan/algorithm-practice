package main

import (
	"bufio"
	"reflect"
	"slices"
	"sort"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	tasks, _, xp, leave := process(reader)
	expect_p := readNum(reader)
	expect_leave := readNNums(reader, len(leave))

	if !reflect.DeepEqual(leave, expect_leave) {
		t.Fatalf("sample expect %v, but got %v", expect_leave, leave)
	}
	// xp 和 expect_p的位置是一样的
	n := len(tasks)
	ps := make([]int, n)
	var x int
	for i := range n {
		ps[i] = tasks[i][2]
		if ps[i] == -1 {
			x = i
		}
	}
	tmp := slices.Clone(ps)
	tmp[x] = expect_p
	sort.Ints(tmp)
	p1 := sort.SearchInts(tmp, expect_p)
	copy(tmp, ps)
	tmp[x] = xp
	sort.Ints(tmp)
	p2 := sort.SearchInts(tmp, xp)
	if p1 != p2 {
		t.Fatalf("sample expect %d, but got %d", expect_p, xp)
	}
}

func TestSample1(t *testing.T) {
	s := `3
4 3 -1
0 2 2
1 3 3
7
4
7 8 4
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3
3 1 2
2 3 3
3 1 -1
4
4
7 6 4
`
	runSample(t, s)
}
