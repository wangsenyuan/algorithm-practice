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
	res := process(reader)
	m := readNum(reader)
	if len(res) != m {
		t.Fatalf("Sample expect %d groups, but got %d", m, len(res))
	}
	slices.SortFunc(res, func(a, b []int) int {
		return a[0] - b[0]
	})

	expect := make([][]int, m)
	for i := range m {
		var k int
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &k) + 1
		expect[i] = make([]int, k)
		for j := range k {
			pos = readInt(s, pos, &expect[i][j]) + 1
		}
	}

	slices.SortFunc(expect, func(a, b []int) int {
		return a[0] - b[0]
	})

	for i := range m {
		sort.Ints(expect[i])
		sort.Ints(res[i])
		if !reflect.DeepEqual(expect[i], res[i]) {
			t.Errorf("Sample expect %v, but got %v", expect[i], res[i])
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 4
1 2
1 3
4 2
4 3
2
2 1 4 
2 2 3`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 1
1 2
1
3 1 2 3 
`)
}
