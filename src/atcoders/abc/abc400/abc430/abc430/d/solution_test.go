package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `10
5 2 7 4 108728325 390529120 597713292 322456626 845148281 812604915
`, []int64{
		10,
		7,
		8,
		8,
		108728326,
		390529121,
		523096670,
		452057486,
		699492475,
		517144218,
	})
}

func TestRBTreeBalancesIncreasingInsertions(t *testing.T) {
	var t0 rbtree
	for i := 0; i < 1000; i++ {
		t0.insert(i)
	}

	var height func(*node) int
	height = func(cur *node) int {
		if cur == nil {
			return 0
		}
		return max(height(cur.ch[0]), height(cur.ch[1])) + 1
	}

	if h := height(t0.root); h > 40 {
		t.Fatalf("tree height = %d, looks unbalanced", h)
	}
}
